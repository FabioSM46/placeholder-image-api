# An API for generating placeholder images

This is a simple API that generates placeholder images with a specified background color and dimensions. It can be used to quickly generate placeholder images for testing or other purposes.

## Usage

To use the API, simply send a GET request to the `/<hexcolor>/widthxheight` endpoint, where `<hexcolor>` is the background color in hex format (e.g. `#000000` for black) and `widthxheight` is the desired width and height of the image (e.g. `100x100`). The API will respond with a JPEG image of the specified dimensions and background color.

For example, to generate a 100x100 black image, you can send a GET request to `http://localhost:8080/#000000/100x100`.

## License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT). 

You are free to use, modify, and distribute this software under the terms of the license. See the `LICENSE` file for more details.