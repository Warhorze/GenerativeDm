from dnd5apy.api import ClassApi

index = 'rogue'
class_client = ClassApi()
print(class_client.api_classes_index_get(index))
#from dnd5epy.api.common_api import CommonApi
#client = CommonApi()
    # getting all possible endpoints for the common API class
#endpoints_names = client.api_get()
#print(endpoints_names)