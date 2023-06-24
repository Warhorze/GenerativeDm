from dnd5epy.api import ClassApi

index = 'figther'
class_client = ClassApi()
print(class_client.api_classes_index_get(index))