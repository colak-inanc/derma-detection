from flask import Flask, request, jsonify
from flask_cors import CORS
import os
from werkzeug.utils import secure_filename
import sys
import json
import logging

# Loglama ayarları
logging.basicConfig(level=logging.DEBUG)
logger = logging.getLogger(__name__)

# Ana dizini Python path'ine ekle
sys.path.append(os.path.dirname(os.path.dirname(os.path.dirname(os.path.abspath(__file__)))))

from src.md.predict_router import route_prediction
from src.md.explain_with_gemini import explain_result

app = Flask(__name__)
CORS(app)

UPLOAD_FOLDER = 'uploads'
ALLOWED_EXTENSIONS = {'png', 'jpg', 'jpeg', 'pdf'}

app.config['UPLOAD_FOLDER'] = UPLOAD_FOLDER

def allowed_file(filename):
    return '.' in filename and filename.rsplit('.', 1)[1].lower() in ALLOWED_EXTENSIONS

@app.route('/process-image', methods=['POST'])
def process_image():
    try:
        logger.debug("Görüntü işleme isteği alındı")
        
        if 'image' not in request.files:
            logger.error("Dosya bulunamadı")
            return jsonify({'error': 'Dosya bulunamadı'}), 400
        
        file = request.files['image']
        if file.filename == '':
            logger.error("Dosya seçilmedi")
            return jsonify({'error': 'Dosya seçilmedi'}), 400
        
        if file and allowed_file(file.filename):
            filename = secure_filename(file.filename)
            filepath = os.path.join(app.config['UPLOAD_FOLDER'], filename)
            
            # Upload dizinini oluştur
            os.makedirs(app.config['UPLOAD_FOLDER'], exist_ok=True)
            
            file.save(filepath)
            logger.debug(f"Dosya kaydedildi: {filepath}")
            
            # Tahmin yap
            sample_input = {
                "image_path": filepath,
                "cbc_data": {
                    "HCT": 50.5,
                    "MCV": 84.7,
                    "PDW": 17.8,
                },
                "symptoms": ["itching", "scaling"]
            }
            
            prediction_results = route_prediction(sample_input)
            logger.debug(f"Tahmin sonucu: {prediction_results}")
            return jsonify(prediction_results)
        
        logger.error("İzin verilmeyen dosya türü")
        return jsonify({'error': 'İzin verilmeyen dosya türü'}), 400
        
    except Exception as e:
        logger.error(f"Hata oluştu: {str(e)}")
        return jsonify({'error': f'İşleme hatası: {str(e)}'}), 500

@app.route('/process-pdf', methods=['POST'])
def process_pdf():
    try:
        logger.debug("PDF işleme isteği alındı")
        
        if 'pdf' not in request.files:
            logger.error("Dosya bulunamadı")
            return jsonify({'error': 'Dosya bulunamadı'}), 400
        
        file = request.files['pdf']
        if file.filename == '':
            logger.error("Dosya seçilmedi")
            return jsonify({'error': 'Dosya seçilmedi'}), 400
        
        if file and allowed_file(file.filename):
            filename = secure_filename(file.filename)
            filepath = os.path.join(app.config['UPLOAD_FOLDER'], filename)
            
            # Upload dizinini oluştur
            os.makedirs(app.config['UPLOAD_FOLDER'], exist_ok=True)
            
            file.save(filepath)
            logger.debug(f"Dosya kaydedildi: {filepath}")
            
            # PDF işleme ve tahmin yapma
            # Bu kısım PDF işleme mantığına göre güncellenecek
            prediction_results = {
                "disease_probability": 0.85,
                "predicted_disease": "Akne",
                "confidence_score": 0.92,
                "recommended_actions": [
                    "Dermatoloğa başvurun",
                    "Reçetesiz akne kremi kullanın",
                    "Yüzünüzü günde iki kez yıkayın"
                ]
            }
            
            logger.debug(f"Tahmin sonucu: {prediction_results}")
            return jsonify(prediction_results)
        
        logger.error("İzin verilmeyen dosya türü")
        return jsonify({'error': 'İzin verilmeyen dosya türü'}), 400
        
    except Exception as e:
        logger.error(f"Hata oluştu: {str(e)}")
        return jsonify({'error': f'İşleme hatası: {str(e)}'}), 500

if __name__ == '__main__':
    os.makedirs(UPLOAD_FOLDER, exist_ok=True)
    logger.info("Python servisi başlatılıyor...")
    app.run(host='0.0.0.0', port=5001, debug=True) 