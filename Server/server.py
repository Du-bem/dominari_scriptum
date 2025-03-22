from flask import Flask, request, jsonify
import requests
import sys, os
sys.path.append(os.path.join(os.path.dirname(__file__), '..', 'Common'))

from encryption import encrypt_json_to_string, decrypt_json_from_string

app = Flask(__name__)

MAIN_SERVER_IP = "127.0.0.0"
MAIN_SERVER_PORT = "2137"
MAIN_SERVER_ENDPOINT = "/api/sendData"

url = f"http://{MAIN_SERVER_IP}:{MAIN_SERVER_PORT}{MAIN_SERVER_ENDPOINT}"


def sendData(data):
    # Set HTTP headers
    headers = {
        'Content-Type': 'application/json'
    }

     # Send the POST request
    response = requests.post(url, headers=headers, json={"data": data})



@app.route('/api/sendData', methods=['POST'])
def receive_data():
    try:
        # Ensure request contains JSON
        data = request.get_json()
        if not data or "data" not in data:
            return jsonify({"error": "Invalid request, 'data' field missing"}), 400

        # Extract encrypted data
        encrypted_data = data["data"]

        decrypted_data = decrypt_json_from_string(encrypted_data)

        sendData(decrypted_data)

        print(f"Received encrypted data: {decrypted_data}")

        return jsonify({"message": "Data received successfully"}), 200

    except Exception as e:
        return jsonify({"error": str(e)}), 500

if __name__ == "__main__":
    app.run(host='0.0.0.0', port=7777, debug=True)
