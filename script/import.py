import csv
from datetime import datetime
from urllib import request, parse
import os
import json

SERVER_ENDPOINT = 'http://localhost:8080'
THIS_FILE = 'script/'

def get_month(month_str):
    months = {
        'ENE': "01",
        'FEB': "02",
        'MAR': "03",
        'ABR': "04",
        'MAY': "05",
        'JUN': "06",
        'JUL': "07",
        'OCT': "08",
        'SEP': "09",
        'OCT': "10",
        'NOV': "11",
        'DIC': "12",
        '1': "01",
        '2': "02",
        '3': "03",
        '4': "04",
        '5': "05",
        '6': "06",
        '7': "07",
        '8': "08",
        '9': "09",
        '10': "10",
        '11': "11",
        '12': "12",
    }

    return months[month_str.upper()]

def parse_date(date_str):
    # Función para convertir la fecha en el formato deseado
    day, month, year = date_str.split('/')
    if (len(day) == 1):
        day = "0" + day

    return "{}-{}-{}".format(year, get_month(month), day)

def read_csv_and_send_data(csv_file_path, company_name):
    movements_data = []

    with open(csv_file_path, 'r') as csvfile:
        csv_reader = csv.reader(csvfile, delimiter=',')
        next(csv_reader, None)

        for row in csv_reader:
            date, shipping_code, pallets, units, code, name, brand, detail, deposit, observations = row

            data_object = {
                "date": parse_date(date),
                "shipping_code": shipping_code,
                "code": code.strip(),
                "name": name.strip(),
                "brand": brand.strip(),
                "detail": detail.strip(),
                "deposit": deposit.strip(),
                "observations": observations.strip()
            }


            if code:
                data_object["code"] = code.strip()
            else:
                data_object["code"] = "No tiene",

            if units:
                data_object["units"] = int(units.strip().replace(',', ''))

            if pallets.strip() != "" and int(pallets.strip()) > 0:
                data_object["units"] = 100000
                data_object["observations"] = "{} {} {}".format(pallets, " pallets", data_object["observations"]).strip()
            
            movements_data.append(data_object)

    url = SERVER_ENDPOINT + "/bulk-create"
    data = {
        "company_name": company_name,
        "movements_data": movements_data
        }
    data_encoded = json.dumps(data).encode('utf-8')
    print(data)
    req = request.Request(url, data=data_encoded, method='POST')
    with request.urlopen(req) as response:
        print(response.status, response.read().decode('utf-8'))

if __name__ == "__main__":
    for file in os.listdir(THIS_FILE):
        if file == 'import.py':
            continue
        
        company_name = file[:-4]
        read_csv_and_send_data(os.path.join(THIS_FILE, file), company_name)