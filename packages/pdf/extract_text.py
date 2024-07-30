# extract_text.py

import sys
import fitz  # PyMuPDF库

def extract_text_from_pdf(pdf_path):
    try:
        print(f"尝试打开PDF文件: {pdf_path}", sys.stderr)
        pdf = fitz.open(pdf_path)
        all_text = []

        for page_num in range(pdf.page_count):
            page = pdf[page_num]
            text = page.get_text()
            all_text.append(text)

        result = "\n".join(all_text)
        print(result)
        return result
    except Exception as e:
        print(f"解析PDF文件时出错: {e}", sys.stderr)
        raise


# if __name__ == "__main__":
#     pdf_path = "packages/pdf/pdf17.pdf"
#     full_text = extract_text_from_pdf(pdf_path)
#     print(full_text)
