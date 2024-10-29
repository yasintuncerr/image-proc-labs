import matplotlib.pyplot as plt
from PIL import Image
import numpy as np

def image_plotter(images, titles=None):
    """
    Plots one or more images with optional titles.
    
    Parameters:
    - images: list of image file paths or PIL Image objects
    - titles: list of titles (strings) corresponding to each image (optional)
    """
    # Ensure images is a list, even if only a single image is provided
    if not isinstance(images, list):
        images = [images]
        
    # Ensure titles is a list, even if only a single title is provided
    if titles is None:
        titles = [''] * len(images)  # Default to empty titles
    elif not isinstance(titles, list):
        titles = [titles]
        
    # Check if the number of titles matches the number of images
    if len(titles) != len(images):
        raise ValueError("The number of titles must match the number of images.")
    
    # Determine the layout for plotting
    plt.figure(figsize=(5 * len(images), 5))
    
    for idx, (img, title) in enumerate(zip(images, titles), 1):
        # Load the image if it is a file path
        if isinstance(img, str):
            img = Image.open(img)
        
        # Convert PIL image to numpy array for plotting
        img_array = np.array(img)
        
        # Plot the image
        plt.subplot(1, len(images), idx)
        plt.imshow(img_array, cmap='gray' if img_array.ndim == 2 else None)
        plt.title(title)
        plt.axis('off')

    plt.tight_layout()
    plt.show()

# Usage example:
# image_plotter(['lenna.bmp', 'airplane.bmp'], ['Lenna', 'Airplane'])
# Or with PIL Images:
# image1 = Image.open('lenna.bmp')
# image2 = Image.open('airplane.bmp')
# image_plotter([image1, image2], ['Lenna', 'Airplane'])
