"""
Reimplementation of Borik's malt command as a Sorik script

Usage: sorik run examples/malt.star --args=image_url=DESIRED_IMAGE_URL,degrees=DEGREES
"""

image_url = get_arg("image_url")
degrees = float(get_arg("degrees"))

image = load_image(image_url)
width, height = image.width, image.height

swirled = swirl(image, degrees)
scaled = liquid_rescale(swirled, width // 2, height // 2)
swirled_rev = swirl(scaled, degrees * -1)
output = liquid_rescale(swirled_rev, width, height)
