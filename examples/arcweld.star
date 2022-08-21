"""
Reimplementation of Borik's arcweld command as a Sorik script

Usage: sorik run examples/arcweld.star --args=image_url=DESIRED_IMAGE_URL
"""

image_url = get_arg("image_url")

image = load_image(image_url)

left_shift_red = evaluate_channel(
    image, ChannelType.CHANNEL_RED, EvaluateOperator.EVAL_OP_LEFT_SHIFT, 1.0
)
contrast_stretched = contrast_stretch(
    left_shift_red,
    0.3,
    0.3,
)
threshold_black = evaluate_channel(
    contrast_stretched,
    ChannelType.CHANNEL_RED,
    EvaluateOperator.EVAL_OP_THRESHOLD_BLACK,
    0.9,
)
sharpened = sharpen(
    threshold_black,
    0.0,
    0.0,
)

scaled_down = liquid_rescale(
    sharpened, sharpened.width // 2, sharpened.height // 3, 1.0, 0.0
)
scaled_up = liquid_rescale(
    scaled_down, scaled_down.width * 2, scaled_down.height * 3, 0.4, 0.0
)

imploded = implode(scaled_up, 0.2)
quantized = quantize(imploded, 8, ColorspaceType.COLORSPACE_RGB, 0, True, False)

output = quantized
