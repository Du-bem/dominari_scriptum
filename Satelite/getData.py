from datetime import datetime, timedelta
from astroquery.jplhorizons import Horizons
import numpy as np
from decimal import Decimal, getcontext
import math
import matplotlib.pyplot as plt
import matplotlib.dates as mdates
import json

cur_time = datetime(2024, 8, 13, 0, 0)

# Set Decimal precision
getcontext().prec = 100  # Set this to the desired level of precision

def date_to_epoch(date):
    """
    Convert a Gregorian date to Julian Date (JD).

    :param date: A datetime.date object representing the Gregorian date.
    :return: The Julian Date corresponding to the given Gregorian date.
    """


    year = date.year
    month = date.month
    day = date.day

    # Convert Gregorian date to Julian Date
    if month <= 2:
        year -= 1
        month += 12
    A = math.floor(year / 100)
    B = 2 - A + math.floor(A / 4)
    jd = (math.floor(365.25 * (year + 4716)) +
          math.floor(30.6001 * (month + 1)) +
          day + B - 1524.5)

    return jd


# Conversion factors with high precision
AU_to_meters = Decimal('1.496e11')  # meters
AU_per_day_to_m_per_s = Decimal('1.496e11') / Decimal('86400')  # meters/second
G = Decimal('6.67430e-11')  # Gravitational constant

i = 1
def load_pos_and_vel(id, time=cur_time):
    global i
    cur_time = datetime(2024, 8, i, 0, 0)
    i += 1
    obj = Horizons(id=id, epochs=date_to_epoch(time)).vectors()
    name = None
    for row in obj:
        name = row[0]
    pos = np.array([Decimal(obj['x'][0]), Decimal(obj['y'][0]), Decimal(obj['z'][0])], dtype=object) * AU_to_meters
    vel = np.array([Decimal(obj['vx'][0]), Decimal(obj['vy'][0]), Decimal(obj['vz'][0])], dtype=object) * AU_per_day_to_m_per_s



    # Convert the numpy arrays and Decimal objects into JSON-serializable lists.
    # Here we convert each Decimal into its string representation to preserve precision.

    pos =  [str(item) for item in pos.tolist()]
    vel =  [str(item) for item in vel.tolist()]

    data_dict = {"name" : name, "time" : str(time), "position" : pos, "velocity" : vel}

    # Convert the dictionary to a JSON string.
    json_output = json.dumps(data_dict, indent=2)

    return json_output
