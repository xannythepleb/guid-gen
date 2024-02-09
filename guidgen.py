import uuid

def generate_guids(num_guids):
  """
  Generates and returns a list of the specified number of random UUIDs (version 4).
  """
  if num_guids <= 0:
    print("Please enter a positive number of GUIDs to generate.")
    return

  guids = []
  for _ in range(num_guids):
    guids.append(str(uuid.uuid4()))

  return guids

# Prompt the user for the number of GUIDs
while True:
  try:
    num_guids = int(input("Welcome to GUID Gen by Xanny. \nHow many GUIDs do you need today? "))
    break
  except ValueError:
    print("Invalid input. Please enter a number above zero.")

# Generate the GUIDs and print them
guids = generate_guids(num_guids)
if guids:
  print("\nGenerated GUIDs: \n")
  for guid in guids:
    print(guid)
