
# Human Brightness Sensitivity Detector

This project is a GUI-based application that allows users to adjust the brightness levels of an image and detect human sensitivity to brightness. It is built using Go and the Fyne GUI framework.

## Features

- **Increase and Decrease Brightness**: Users can increment or decrement the brightness of the image step by step.
- **Set Step**: Users can adjust the step size for brightness changes.
- **Save Brightness Levels**: Users can save different brightness levels and view them in the UI.
- **Calculate Human Sensitivity**: The application calculates the human sensitivity based on saved brightness levels.
- **Image Refresh**: The image refreshes as brightness levels change.

## Technologies Used

- **Go**: The application is written in Go.
- **Fyne**: A Go-based graphical user interface library.
- **Go Generics**: Used for brightness adjustments and calculations.

## Installation

1. **Install Go**: Make sure Go is installed. If not, install it from [here](https://golang.org/dl/).

2. **Clone the Project**:
   ```bash
   git clone https://github.com/yasintuncerr/image-proc-labs.git
   cd image-proc-labs/Chapters/ch1
   ```

3. **Install Dependencies**:
   Install the required dependencies:
   ```bash
   go mod tidy
   ```

4. **Run the Application**:
   Run the application using the following command:
   ```bash
   go run .
   ```

## Usage

1. **Adjust Brightness**: Click the `Increase` button to increment brightness, or `Decrease` to reduce it. The current level will be displayed, and saved levels will be updated.
2. **Set Step**: Use the `Set Step` button to change the step size for brightness adjustments.
3. **Save Brightness Level**: Click the `Save` button to save the current brightness level.
4. **Calculate Sensitivity**: Click the `Calculate Human Sensitivity` button to calculate brightness sensitivity based on saved levels.

## File Structure

- `app.go`: Entry point of the application. Defines all UI components and manages interactions.
- `brightness.go`: Contains functions for brightness adjustment on the image.
- `controller.go`: Manages application logic.
- `model.go`: Holds the image and brightness-related data.
- `view.go`: Handles UI updates and image processing.

## Contributing

If you'd like to contribute, feel free to submit a pull request or open an issue to discuss improvements or feature requests.

## License

This project is licensed under the MIT License. See the `LICENSE` file for more details.
