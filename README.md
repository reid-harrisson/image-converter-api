# Image Converter API

This is a simple backend API built with [Fiber](https://gofiber.io/), a web framework for Golang, designed to remove the background from PNG or JPG images. The API provides a single endpoint to process images and return the result with the background removed.

## Features

- Remove background from PNG and JPG images.
- Fast and efficient processing using Golang and Fiber.
- Simple and easy-to-use API.

## Getting Started

### Prerequisites

- Go 1.16 or later
- [Fiber](https://gofiber.io/) framework

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/image-converter-api.git
   cd image-converter-api
   ```

2. Install dependencies:

   ```bash
   go get
   ```

3. Run the server:

   ```bash
   go run main.go
   ```

### Usage

The API provides a single endpoint to remove the background from an image.

#### Endpoint

- **POST** `/remove-background`

#### Request

- **Headers**: `Content-Type: multipart/form-data`
- **Body**: Form-data with a file field named `image` containing the PNG or JPG file.

#### Response

- **Success**: Returns the image with the background removed.
- **Error**: Returns an error message if the processing fails.

### Example

You can use `curl` to test the API:

```bash
curl -X POST http://localhost:3000/image/convert?fore_color=#FFFFFF&back_color=#000000 \
  -F "file=@/path/to/your/image.png" \
  -o output.png
```

### Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

### License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

### Acknowledgments

- Thanks to the Fiber team for their excellent framework.
- Inspired by various open-source image processing tools.
