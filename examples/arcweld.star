"""
Reimplementation of Borik's arcweld command as a Sorik script

Usage: sorik run examples/arcweld.star --args=image_url=DESIRED_IMAGE_URL
"""

image_url = get_arg("image_url")
image = load_image(image_url)

image = evaluate_channel(
    image, ChannelType.CHANNEL_RED, EvaluateOperator.EVAL_OP_LEFT_SHIFT, 1.0
)
image = contrast_stretch(
    image,
    0.3,
    0.3,
)
image = evaluate_channel(
    image,
    ChannelType.CHANNEL_RED,
    EvaluateOperator.EVAL_OP_THRESHOLD_BLACK,
    0.9,
)
image = sharpen(
    image,
    0.0,
    0.0,
)
image = liquid_rescale(image, image.width // 2, image.height // 3, 1.0, 0.0)
image = liquid_rescale(image, image.width * 2, image.height * 3, 0.4, 0.0)
image = implode(image, 0.2)
image = quantize(image, 8, ColorspaceType.COLORSPACE_RGB, 0, True, False)

output = image
