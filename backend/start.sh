#!/bin/bash

# Python servisini başlat
cd python_service
python3 -m venv venv
source venv/bin/activate
pip install -r requirements.txt
python app.py &
PYTHON_PID=$!

# Go servisini başlat
cd ..
go run main.go &
GO_PID=$!

# Servisleri durdurmak için Ctrl+C'yi yakala
trap "kill $PYTHON_PID $GO_PID; exit" INT

# Servislerin çalışmasını bekle
wait 