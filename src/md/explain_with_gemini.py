import google.generativeai as genai

def explain_result(prompt):
    """
    Gemini API kullanarak tahmin sonuçlarını açıklar.
    
    Args:
        prompt (str): Açıklanacak tahmin sonuçları
    
    Returns:
        str: Açıklama metni
    """
    # Örnek açıklama
    return """
    Tahmin sonuçlarına göre, yüksek olasılıkla (%85) akne durumu söz konusu. 
    Bu tahmin, görüntü analizi ve kan testi sonuçlarının birleşiminden elde edilmiştir.
    
    Önerilen aksiyonlar:
    1. En kısa sürede bir dermatoloğa başvurmanız önerilir
    2. Reçetesiz akne kremi kullanabilirsiniz
    3. Cilt temizliğine özen göstermeniz önemli
    
    Not: Bu bir ön değerlendirmedir ve kesin teşhis için mutlaka bir uzmana başvurulmalıdır.
    """ 