import csv
from datetime import datetime
from urllib import request, parse
import os

YEAR = datetime.now().year
SERVER_ENDPOINT = 'localhost:8080'
THIS_FILE = 'script/'

def get_month(month_str):
    months = {
        'ENE': 1,
        'FEB': 2,
        'MAR': 3,
        'ABR': 4,
        'MAY': 5,
        'JUN': 6,
        'JUL': 7,
        'OCT': 8,
        'SEP': 9,
        'OCT': 10,
        'NOV': 11,
        'DIC': 12,
    }

    return months[month_str.upper()]

def parse_date(date_str):
    # Función para convertir la fecha en el formato deseado
    day, month = date_str.split('-')

    return datetime(YEAR, get_month(month), int(day))

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

            if pallets:
                data_object["pallets"] = int(pallets.strip())
            if units:
                data_object["units"] = int(units.strip().replace(',', ''))
            
            movements_data.append(data_object)

    url = SERVER_ENDPOINT + "/bulk"
    data = {
        "company_name": company_name,
        "movements_data": movements_data
        }
    
    data = parse.urlencode(data).encode('utf-8')
    req = request.Request(url, data=data, method='POST')
    with request.urlopen(req) as response:
        print(response.status, response.read().decode('utf-8'))

if __name__ == "__main__":
    for file in os.listdir(THIS_FILE):
        if file == 'import.py':
            continue
        
        company_name = file[:-4]
        read_csv_and_send_data(os.path.join(THIS_FILE, file), company_name)