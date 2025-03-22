import json
import os
import base64
import hashlib
from cryptography.hazmat.primitives.ciphers import Cipher, algorithms, modes
from cryptography.hazmat.primitives.kdf.pbkdf2 import PBKDF2HMAC
from cryptography.hazmat.primitives import hashes
from cryptography.hazmat.backends import default_backend

# Constants
SECRET_KEY_FILE = "secret.key"
SALT = b"my_secure_salt_value"  # Use a securely generated salt
IV_SIZE = 16  # AES block size
KEY_SIZE = 32  # AES-256 requires a 32-byte key


def generate_key(passphrase):
    """Derive a 256-bit key from a passphrase and save it."""
    kdf = PBKDF2HMAC(
        algorithm=hashes.SHA256(),
        length=KEY_SIZE,
        salt=SALT,
        iterations=100000,  # Adjust for security vs. performance
        backend=default_backend(),
    )
    key = kdf.derive(passphrase.encode())  # Derive a 256-bit key from passphrase

    # Save the key to a file
    with open(SECRET_KEY_FILE, "wb") as f:
        f.write(key)


def load_key():
    """Load the AES key from the file."""
    with open(SECRET_KEY_FILE, "rb") as f:
        return f.read()


def encrypt_json_to_string(json_data):
    """Encrypt JSON and return a Base64-encoded string."""
    key = load_key()

    # Convert JSON object to a string
    json_string = json.dumps(json_data)

    # Generate a random IV
    iv = os.urandom(IV_SIZE)

    # Pad data to be a multiple of AES block size (16 bytes)
    pad_len = 16 - (len(json_string) % 16)
    padded_data = json_string + (chr(pad_len) * pad_len)

    # Encrypt using AES-256-CBC
    cipher = Cipher(algorithms.AES(key), modes.CBC(iv), backend=default_backend())
    encryptor = cipher.encryptor()
    encrypted_data = encryptor.update(padded_data.encode()) + encryptor.finalize()

    # Combine IV and encrypted data, then encode as Base64
    encrypted_payload = base64.b64encode(iv + encrypted_data).decode()

    return encrypted_payload  # This string can be sent over the internet


def decrypt_json_from_string(encrypted_string):
    """Decrypt a Base64-encoded AES-encrypted string and return JSON data."""
    key = load_key()

    # Decode Base64
    encrypted_payload = base64.b64decode(encrypted_string)

    # Extract IV and encrypted data
    iv = encrypted_payload[:IV_SIZE]
    encrypted_data = encrypted_payload[IV_SIZE:]

    # Decrypt using AES-256-CBC
    cipher = Cipher(algorithms.AES(key), modes.CBC(iv), backend=default_backend())
    decryptor = cipher.decryptor()
    decrypted_padded_data = decryptor.update(encrypted_data) + decryptor.finalize()

    # Remove padding
    pad_len = ord(decrypted_padded_data[-1:])
    decrypted_data = decrypted_padded_data[:-pad_len].decode()

    # Convert back to JSON
    return json.loads(decrypted_data)


generate_key("passphrase")  # Generate a strong key from the passphrase

# Sample JSON data
