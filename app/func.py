from parliament import Context, event


@event
def main(context: Context):
    """
    Function template
    The context parameter contains the Flask request object and any
    CloudEvent received with the request.
    """

    # Add your business logic here
    print(context.cloud_event.data)

    # The return value here will be applied as the data attribute
    # of a CloudEvent returned to the function invoker
    return context.cloud_event.data
