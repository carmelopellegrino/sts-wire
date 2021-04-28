package core

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/DODAS-TS/sts-wire/pkg/rclone"
	"github.com/rs/zerolog/log"
)

const (
	numCheckAttempts = 10
	checkExeWait     = 1 * time.Second
	waitFileBusy     = 1 * time.Second
	closeChannelWait = 2 * time.Second
)

func CacheDir() (string, error) {
	cacheDir, errCacheDir := os.UserCacheDir()
	if errCacheDir != nil {
		log.Err(errCacheDir).Msg("Caching dir not available")

		return "", fmt.Errorf("CacheDir %w", errCacheDir)
	}

	return filepath.Join(cacheDir, "sts-wire"), nil
}

func ExePath() (string, error) {
	cacheDir, errCacheDir := CacheDir()
	if errCacheDir != nil {
		return "", errCacheDir
	}

	return filepath.Join(cacheDir, rcloneExeString), nil
}

func CheckExeFile(rcloneFile string, originalData []byte) error { //nolint:funlen
	var rcloneExeFile bytes.Buffer

	rcloneFileReader, errOpen := os.Open(rcloneFile)
	if errOpen != nil {
		log.Err(errOpen).Msg("Cannot open rclone executable in cache dir")

		return fmt.Errorf("prepare rclone %w", errOpen)
	}

	defer rcloneFileReader.Close()

	// Check if it has executable permissions
	if runtime.GOOS != "windows" {
		rcloneInfo, err := rcloneFileReader.Stat()
		if err != nil {
			return fmt.Errorf("cannot have rclone file stat")
		}

		if rcloneInfo.Mode() != os.FileMode(exeFileMode) {
			return fmt.Errorf("rclone is not an executable")
		}
	}

	// Verify checksum
	_, errRead := rcloneExeFile.ReadFrom(rcloneFileReader)
	if errRead != nil {
		log.Err(errRead).Msg("Cannot read rclone executable in cache dir")

		return fmt.Errorf("prepare rclone %w", errRead)
	}

	curExe := md5.New() //nolint:gosec

	_, errWriteCurExe := io.Copy(curExe, bytes.NewReader(rcloneExeFile.Bytes()))
	if errWriteCurExe != nil {
		log.Err(errWriteCurExe).Msg("Cannot calculate md5 for rclone executable in cache dir")

		return fmt.Errorf("prepare rclone %w", errWriteCurExe)
	}

	rcloneExe := md5.New() //nolint:gosec

	_, errWriteRcloneExe := io.Copy(rcloneExe, bytes.NewReader(originalData))
	if errWriteRcloneExe != nil {
		log.Err(errWriteRcloneExe).Msg("Cannot calculate md5 for original rclone")

		return fmt.Errorf("prepare rclone %w", errWriteRcloneExe)
	}

	rcloneMd5 := curExe.Sum(nil)
	originalMd5 := rcloneExe.Sum(nil)
	log.Debug().Str("rcloneMd5",
		hex.EncodeToString(rcloneMd5)).Str("originalMd5",
		hex.EncodeToString(originalMd5)).Msg("rclone - executable checksum")

	if !bytes.Equal(rcloneMd5, originalMd5) {
		log.Err(nil).Msg("checksum not equal")

		return fmt.Errorf("checksum not equal")
	}

	return nil
}

func PrepareRclone() error { // nolint: funlen,gocognit,cyclop
	baseDir, errCacheDir := CacheDir()
	if errCacheDir != nil {
		return errCacheDir
	}

	log.Debug().Str("basedir", baseDir).Msg("rclone")

	rcloneFile, errExePath := ExePath()
	if errExePath != nil {
		return errExePath
	}

	log.Debug().Str("rcloneFile", rcloneFile).Msg("rclone")
	log.Debug().Msg("rclone - get asset data")

	data := rclone.Executable
	log.Debug().Int("assetLen", len(data)).Msg("rclone")

	log.Debug().Msg("rclone - create executable base directories")

	errMkdir := os.MkdirAll(baseDir, os.ModePerm)
	if errMkdir != nil && !os.IsExist(errMkdir) {
		log.Err(errMkdir).Msg("Cannot create rclone cache dir")

		return fmt.Errorf("prepare rclone %w", errMkdir)
	}

	_, errStat := os.Stat(rcloneFile)
	log.Debug().Bool("missingExe", os.IsNotExist(errStat)).Msg("rclone")

	if errStat != nil && os.IsNotExist(errStat) { //nolint:nestif
		log.Debug().Msg("rclone - creating executable")

		rcloneExeFile, errCreate := os.OpenFile(rcloneFile, os.O_RDWR|os.O_CREATE|os.O_SYNC, fileMode)
		if errCreate != nil {
			log.Err(errCreate).Msg("rclone - creating executable - Cannot open with write permission the rclone executable in cache dir")

			// file busy
			if strings.Contains(errCreate.Error(), "file busy") {
				var errCheck error

				for attempt := 0; attempt < numCheckAttempts; attempt++ {
					log.Debug().Int("attempt", attempt).Msg("rclone - verify executable")

					errCheck = CheckExeFile(rcloneFile, data)
					if errCheck != nil {
						log.Err(errCheck).Int("attempt", attempt).Msg("rclone - verify executable - Cannot verify rclone executable in cache dir when file is busy")
					} else {
						break
					}

					time.Sleep(waitFileBusy)
				}

				if errCheck != nil {
					return errCheck
				}

				return nil
			}

			log.Err(errCreate).Msg("rclone - creating executable - Cannot create rclone executable in cache dir")

			return fmt.Errorf("prepare rclone %w", errCreate)
		}

		buff := bytes.NewReader(data)
		writtenData, errWrite := io.Copy(rcloneExeFile, buff)

		log.Debug().Int64("writtenData", writtenData).Msg("rclone")

		if errWrite != nil {
			log.Err(errWrite).Msg("rclone - creating executable - Cannot write rclone executable in cache dir")

			return fmt.Errorf("prepare rclone %w", errWrite)
		}

		rcloneExeFile.Close()

		log.Debug().Msg("rclone - change executable mod")

		errChmod := os.Chmod(rcloneFile, os.FileMode(exeFileMode))
		if errChmod != nil {
			log.Err(errChmod).Msg("rclone - creating executable - Cannot make rclone an executable in cache dir")

			return fmt.Errorf("prepare rclone %w", errChmod)
		}
	}

	var errCheck error

	for attempt := 0; attempt < numCheckAttempts; attempt++ {
		log.Debug().Int("attempt", attempt).Msg("rclone - verify executable")

		errCheck = CheckExeFile(rcloneFile, data)
		if errCheck != nil {
			log.Err(errCheck).Int("attempt", attempt).Msg("Cannot verify rclone executable in cache dir")
		} else {
			break
		}

		time.Sleep(checkExeWait)
	}

	if errCheck != nil {
		return fmt.Errorf("prepare rclone %w", errCheck)
	}

	return nil
}

func MountVolume(instance string, remotePath string, localPath string, configPath string, readOnly bool, noModtime bool, newFlags string) (*exec.Cmd, chan error, string, error) { // nolint: funlen,gocognit,lll,cyclop
	log.Debug().Str("action", "prepare rclone").Msg("rclone - mount")

	errPrepare := PrepareRclone()
	if errPrepare != nil {
		log.Err(errPrepare).Msg("rclone - mount")

		return nil, nil, "", errPrepare
	}

	log.Debug().Str("action", "get file path").Msg("rclone - mount")

	rcloneFile, errExePath := ExePath()
	if errExePath != nil {
		return nil, nil, "", errExePath
	}

	log.Debug().Str("action", "make local dir").Msg("rclone - mount")

	if runtime.GOOS != "windows" {
		_, errLocalPath := os.Stat(localPath)
		if os.IsNotExist(errLocalPath) {
			errMkdir := os.MkdirAll(localPath, os.ModePerm)
			if errMkdir != nil {
				panic(errMkdir)
			}
		}
	}

	conf := fmt.Sprintf("%s:%s", instance, remotePath)

	log.Debug().Str("action", "prepare mounting points").Msg("rclone - mount")

	configPathAbs, errConfigPath := filepath.Abs(filepath.Join(configPath, "/rclone.conf"))
	if errConfigPath != nil {
		log.Err(errConfigPath).Msg("rclone - mount")

		return nil, nil, "", fmt.Errorf("rclone config abs: %w", errConfigPath)
	}

	logPath, errLogPath := filepath.Abs(filepath.Join(configPath, "/rclone.log"))
	if errLogPath != nil {
		log.Err(errLogPath).Msg("rclone - mount")

		return nil, nil, "", fmt.Errorf("rclone log abs: %w", errLogPath)
	}

	localPathAbs, errLocalPath := filepath.Abs(localPath)
	if errLocalPath != nil {
		log.Err(errLocalPath).Msg("rclone - mount")

		return nil, nil, "", fmt.Errorf("local path abs: %w", errLocalPath)
	}

	commandArgs := strings.Builder{}

	commandArgs.WriteString("--config")
	commandArgs.WriteRune(' ')
	commandArgs.WriteString(configPathAbs)
	commandArgs.WriteRune(' ')
	// commandArgs.WriteString(// "--daemon")
	// commandArgs.WriteRune(' ')
	commandArgs.WriteString("--log-file")
	commandArgs.WriteRune(' ')
	commandArgs.WriteString(logPath)
	commandArgs.WriteRune(' ')
	commandArgs.WriteString("--log-level")
	commandArgs.WriteRune(' ')
	commandArgs.WriteString("DEBUG")
	commandArgs.WriteRune(' ')
	// commandArgs.WriteString("--use-json-log")
	// commandArgs.WriteRune(' ')
	commandArgs.WriteString("--no-check-certificate")
	commandArgs.WriteRune(' ')
	commandArgs.WriteString("--local-no-sparse")
	commandArgs.WriteRune(' ')
	commandArgs.WriteString("--fast-list")
	commandArgs.WriteRune(' ')
	commandArgs.WriteString("--s3-disable-http2")
	commandArgs.WriteRune(' ')
	commandArgs.WriteString("mount")
	commandArgs.WriteRune(' ')
	commandArgs.WriteString(conf)
	commandArgs.WriteRune(' ')
	commandArgs.WriteString(localPathAbs)

	commandFlags := strings.Builder{}
	commandFlags.WriteString("--debug-fuse")
	commandFlags.WriteRune(' ')
	commandFlags.WriteString("--write-back-cache")
	commandFlags.WriteRune(' ')
	commandFlags.WriteString("--attr-timeout")
	commandFlags.WriteRune(' ')
	commandFlags.WriteString("20m")
	commandFlags.WriteRune(' ')
	commandFlags.WriteString("--dir-cache-time")
	commandFlags.WriteRune(' ')
	commandFlags.WriteString("20m")
	commandFlags.WriteRune(' ')
	commandFlags.WriteString("--vfs-cache-poll-interval")
	commandFlags.WriteRune(' ')
	commandFlags.WriteString("20m")
	commandFlags.WriteRune(' ')
	commandFlags.WriteString("--vfs-write-back")
	commandFlags.WriteRune(' ')
	commandFlags.WriteString("30s")
	commandFlags.WriteRune(' ')
	commandFlags.WriteString("--vfs-write-wait")
	commandFlags.WriteRune(' ')
	commandFlags.WriteString("10s")
	commandFlags.WriteRune(' ')
	commandFlags.WriteString("--vfs-read-wait")
	commandFlags.WriteRune(' ')
	commandFlags.WriteString("30s")
	commandFlags.WriteRune(' ')
	commandFlags.WriteString("--vfs-cache-max-size")
	commandFlags.WriteRune(' ')
	commandFlags.WriteString("4G")
	commandFlags.WriteRune(' ')
	commandFlags.WriteString("--vfs-cache-mode")
	commandFlags.WriteRune(' ')
	commandFlags.WriteString("writes")

	if noModtime {
		commandFlags.WriteRune(' ')
		commandFlags.WriteString("--no-modtime")
	}

	if readOnly {
		commandFlags.WriteRune(' ')
		commandFlags.WriteString("--read-only")
	}

	commandArgs.WriteRune(' ')

	if newFlags != "" {
		commandArgs.WriteString(newFlags)
	} else {
		commandArgs.WriteString(commandFlags.String())
	}

	log.Debug().Str("command",
		rcloneFile).Str("args",
		commandArgs.String(),
	).Msg("rclone - mount")

	rcloneCmd := exec.Command(rcloneFile, strings.Split(commandArgs.String(), " ")...)

	log.Debug().Str("action", "start rclone").Msg("rclone - mount")

	if errStart := rcloneCmd.Start(); errStart != nil {
		panic(errStart)
	}

	cmdErr := make(chan error)

	cmdErrorCheck := func() {
		defer close(cmdErr)

		procState, errWait := rcloneCmd.Process.Wait()

		// Wait for main server to close channel if User pressed Ctrl+C
		time.Sleep(closeChannelWait)

		openChannel := true
		// Check if channel still is open
		select {
		case _, ok := <-cmdErr:
			if !ok {
				openChannel = false
			}
		default:
		}

		if procState != nil {
			// ref: https://rclone.org/docs/#exit-code
			switch exitCode := procState.ExitCode(); exitCode {
			case 0:
				log.Debug().Int("exitCode",
					exitCode).Msg("success")
			case 1:
				log.Error().Int("exitCode",
					exitCode).Msg("Syntax or usage error")
			case 2:
				log.Error().Int("exitCode",
					exitCode).Msg("Error not otherwise categorised")
			case 3:
				log.Error().Int("exitCode",
					exitCode).Msg("Directory not found")
			case 4:
				log.Error().Int("exitCode",
					exitCode).Msg("File not found")
			case 5:
				log.Error().Int("exitCode",
					exitCode).Msg("Temporary error (one that more retries might fix) (Retry errors)")
			case 6:
				log.Error().Int("exitCode",
					exitCode).Msg("Less serious errors (like 461 errors from dropbox) (NoRetry errors)")
			case 7:
				log.Error().Int("exitCode",
					exitCode).Msg("Fatal error (one that more retries won't fix, like account suspended) (Fatal errors)")
			case 8:
				log.Error().Int("exitCode",
					exitCode).Msg("Transfer exceeded - limit set by --max-transfer reached")
			case 9:
				log.Error().Int("exitCode",
					exitCode).Msg("Operation successful, but no files transferred")
			default:
				log.Error().Int("exitCode",
					exitCode).Msg("error on exit")
			}
		}

		if openChannel { //nolint:nestif
			// rclone exited with errors
			if errWait != nil {
				cmdErr <- errWait
			}
		} else {
			if errWait == nil {
				// rclone not exited after user pressed Ctrl+C
				if !procState.Exited() {
					panic("rclone termination error")
				}
			} else if errWait.Error() != "wait: no child processes" {
				// rclone exited for unknown reason
				panic(errWait)
			}
		}
	}
	go cmdErrorCheck()

	return rcloneCmd, cmdErr, logPath, nil
}

type RcloneLogErrorMsg struct {
	LineNumber int
	Str        string
	LookupFile string
}

func RcloneLogRotate(logPath string) {
	readFile, err := os.Open(logPath)
	if err != nil {
		log.Err(err).Str("logPath", logPath).Msg("failed to open log file for rotation")
	}

	defer readFile.Close()

	logDir := filepath.Dir(logPath)

	allFiles, err := ioutil.ReadDir(logDir)
	if err != nil {
		log.Err(err).Str("logPath", logPath).Msg("failed to search other log files for rotation")
	}

	lastLogNum := -1

	for _, file := range allFiles {
		if strings.Contains(file.Name(), "rclone") && strings.Contains(file.Name(), ".log") { //nolint:nestif
			curNum := strings.Replace(file.Name(), "rclone", "", 1)
			curNum = strings.Replace(curNum, ".log", "", 1)
			curNum = strings.Replace(curNum, ".gz", "", 1)

			if len(curNum) == 0 {
				lastLogNum = 0
			} else {
				tmpNum, err := strconv.ParseInt(curNum, 10, 32)
				if err != nil {
					log.Err(err).Str("logPath", logPath).Msg("failed to convert log file num for rotation")
				}

				if int(tmpNum) > lastLogNum {
					lastLogNum = int(tmpNum)
				}
			}

			log.Debug().Str("curNum", curNum).Int("lastLogNum", lastLogNum).Msg("search other log files for rotation")
		}
	}

	lastLogNum++

	log.Debug().Int("lastLogNum", lastLogNum).Msg("next log file for rotation")

	newLogFileName := filepath.Join(logDir, fmt.Sprintf("rclone%d.log.gz", lastLogNum))

	newLogFile, err := os.Create(newLogFileName)
	if err != nil {
		log.Err(err).Str("logPath", logPath).Msg("failed to create log file for rotation")
	}

	defer newLogFile.Close()

	writer, err := gzip.NewWriterLevel(newLogFile, gzip.BestCompression)
	if err != nil {
		log.Err(err).Str("logPath", logPath).Msg("failed to create gzip log file writer for rotation")
	}

	defer writer.Close()

	_, err = io.Copy(writer, readFile)
	if err != nil {
		log.Err(err).Str("logPath", logPath).Msg("failed to copy log file for rotation")
	}

	err = os.Truncate(logPath, 0)
	if err != nil {
		log.Err(err).Str("logPath", logPath).Msg("failed truncate log file for rotation")
	}

	log.Debug().Str("logPath", logPath).Int("numRotation", lastLogNum).Msg("log file rotated")
}

func RcloneLogErrors(logPath string, fromLine int) chan RcloneLogErrorMsg { //nolint:funlen,cyclop
	outErrors := make(chan RcloneLogErrorMsg)

	go func() {
		defer close(outErrors)

		latestErrors := make([]RcloneLogErrorMsg, 0)

		readFile, err := os.Open(logPath)
		if err != nil {
			log.Err(err).Str("logPath", logPath).Msg("failed to open log file")
		}

		defer readFile.Close()

		fileInfo, err := readFile.Stat()
		if err != nil {
			log.Err(err).Str("logPath", logPath).Msg("failed to get stats of the log file")
		}

		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)

		lineNum := 0
		curLookupFile := ""

		for fileScanner.Scan() {
			if lineNum >= fromLine {
				curLine := fileScanner.Text()

				switch {
				case strings.Contains(curLine, "INFO") && strings.Contains(curLine, "Exiting..."):
					latestErrors = make([]RcloneLogErrorMsg, 0)
				case strings.Contains(curLine, "error"), strings.Contains(curLine, "ERROR"):
					latestErrors = append(latestErrors, RcloneLogErrorMsg{
						LineNumber: lineNum,
						Str:        curLine,
						LookupFile: curLookupFile,
					})
				case strings.Contains(curLine, "LOOKUP /"):
					parts := strings.Split(curLine, " ")
					filename := parts[1][1:]
					curLookupFile = filename

					log.Debug().Str("lookup", curLookupFile).Msg("lookup")
				}
			}

			lineNum++
		}

		if fileInfo.Size() >= oneMB*logMaxSizeMB {
			go RcloneLogRotate(logPath)

			latestErrors = append(latestErrors, RcloneLogErrorMsg{
				LineNumber: 0,
				Str:        "LOGROTATE",
				LookupFile: "",
			})
		}

		for _, foundErr := range latestErrors {
			outErrors <- foundErr
		}
	}()

	return outErrors
}
