from getData import load_pos_and_vel
import sys, os
import time
sys.path.append(os.path.join(os.path.dirname(__file__), '..', 'Common'))

from encryption import encrypt_json_to_string, decrypt_json_from_string


# print(decrypt_json_from_string())

import requests

# Define constants
IP = "10.44.63.237"
PORT = "7777"
ENDPOINT = "/api/sendData"

# Construct the target URL
url = f"http://{IP}:{PORT}{ENDPOINT}"


# Set HTTP headers
headers = {
    'Content-Type': 'application/json'
}

# Continuous data sending loop
while True:
    try:
        time.sleep(1)

        # Load and encrypt data
        raw_data = load_pos_and_vel(399)
        encrypted_data = encrypt_json_to_string(raw_data)

        # Send the POST request
        response = requests.post(url, headers=headers, json={"data": encrypted_data})

        # Print response details
        print(f"Response Code: {response.status_code}")
        print("Response Body:")
        print(response.text)

    except requests.exceptions.RequestException as e:
        print(f"Error sending data: {e}")
