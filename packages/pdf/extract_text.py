import sys
import fitz  # PyMuPDF库
import io

# 强制设置标准输出为UTF-8编码
sys.stdout = io.TextIOWrapper(sys.stdout.buffer, 'utf-8')

def extract_text_from_pdf(pdf_path):
    try:
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

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: python extract_text.py <pdf_path>", sys.stderr)
        # sys.exit(1)
    pdf_path = sys.argv[1]
    extract_text_from_pdf(pdf_path)
