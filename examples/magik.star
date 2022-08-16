"""
Reimplementation of Borik's magik command as a Sorik script

Usage: sorik run examples/magik.star --args=image_url=DESIRED_IMAGE_URL,scale=INTENSITY
"""

image_url = get_arg("image_url")
scale = float(get_arg("scale"))

image = load_image(image_url)

width, height = image.width, image.height

scaled = liquid_rescale(image, int(width / 2), int(height / 2), deltax=scale)
output = liquid_rescale(scaled, width, height, deltax=scale)
