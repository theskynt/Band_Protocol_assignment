import requests
import time

class TransactionClient:
    def __init__(self, base_url):
        self.base_url = base_url

    def broadcast_transaction(self, symbol, price, timestamp):
        """
        ส่งคำขอ POST เพื่อส่งธุรกรรมไปยังเซิร์ฟเวอร์

        Args:
            symbol (str): ชื่อธุรกรรม
            price (int): ราคาของธุรกรรม
            timestamp (int): ประทับเวลาของการรับราคา

        Returns:
            str: ค่า hash ของธุรกรรมที่ถูกสร้าง
        """
        payload = {
            "symbol": symbol,
            "price": price,
            "timestamp": timestamp
        }
        response = requests.post(f"{self.base_url}/broadcast", json=payload)
        return response.json().get("tx_hash")

    def check_transaction_status(self, tx_hash):
        """
        ส่งคำขอ GET เพื่อตรวจสอบสถานะของธุรกรรม

        Args:
            tx_hash (str): ค่า hash ของธุรกรรมที่ต้องการตรวจสอบสถานะ

        Returns:
            str: สถานะของธุรกรรม ("CONFIRMED", "FAILED", "PENDING", "DNE")
        """
        response = requests.get(f"{self.base_url}/check/{tx_hash}")
        json_data = response.json()
        return json_data.get("tx_status", "UNKNOWN")

def main():
    # เริ่มต้น client ด้วย URL พื้นฐานของเซิร์ฟเวอร์
    client = TransactionClient("https://mock-node-wgqbnxruha-as.a.run.app")

    # ตัวอย่างการส่งธุรกรรม
    symbol = "ETH"
    price = 4500
    timestamp = 1678912345
    tx_hash = client.broadcast_transaction(symbol, price, timestamp)
    print("Transaction hash:", tx_hash)

    # ตรวจสอบสถานะของธุรกรรม
    while True:
        status = client.check_transaction_status(tx_hash)
        print("Transaction status:", status)
        if status == "CONFIRMED" or status == "FAILED" or status == "DNE":
            break
        time.sleep(5)  # ตรวจสอบสถานะทุก ๆ 5 วินาที

if __name__ == "__main__":
    main()
