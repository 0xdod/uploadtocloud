# UploadToCloud CLI Tool

The UploadToCloud CLI Tool is a convenient command-line application developed in Go, designed to simplify the process of uploading local files to your preferred cloud storage service directly from your terminal. This tool streamlines file sharing and backup tasks by eliminating the need for manual uploads through web interfaces.

## Features

- **Effortless Uploads:** Easily upload files and directories to your chosen cloud storage service using simple command-line instructions.

- **Cross-Platform Compatibility:** UploadToCloud runs smoothly on various platforms, including Windows, macOS, and Linux.

- **Multi-Cloud Support:** Seamlessly connect to popular cloud storage providers such as Dropbox, Google Drive, and Amazon S3.

- **Upload Progress Tracking:** Stay informed about upload progress through interactive progress indicators.

- **Configuration Options:** Configure cloud storage credentials and default settings using an intuitive configuration file.

- **Flexible File Handling:** Upload individual files or entire directories, with the option to include subdirectories.

## Installation

You can install UploadToCloud using the Go package manager. Open your terminal and execute:

```bash
go get -u github.com/0xdod/uploadtocloud
```

## Usage

1. Configure UploadToCloud with your cloud storage credentials using the `.config.sample.yaml` file as template for the `.config.yaml` config file

2. Upload a local file to your cloud storage service with the following command:

```bash
uploadtocloud upload /path/to/local/file.txt
```

3. Track upload progress through the interactive progress bar displayed in the terminal.

## Configuration

After installing UploadToCloud, you need to configure it with your cloud storage credentials.

## Supported Cloud Services

UploadToCloud currently supports the following cloud storage services:
- [x] Amazon S3

Additional cloud service support is planned for future updates.

## TODO

Implement everything I just wrote in this README, this thing doesn't work yet.

## License

This project is licensed under the MIT License. Refer to the [LICENSE](LICENSE) file for details.

---

Simplify your cloud file uploads with the UploadToCloud CLI Tool, built with the power of Go. If you encounter any challenges or have ideas for improvement, please feel free to open an issue on the GitHub repository.

Enjoy hassle-free uploading! 🚀
