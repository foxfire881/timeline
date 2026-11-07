# import requests

# x = requests.post("http://10.0.0.3:8080/api/wx/send", json={"msg": "999"}).content.decode("utf-8")
# print(x)

import json,sqlite3
conn = sqlite3.connect("../timeline.db")

def date_to_utc_timestamp(date_str: str) -> int:
    """
    纯Python内置：公元前/公元日期 → UTC Unix 时间戳(秒)
    输出一定是 00:00:00 UTC
    无任何第三方库
    """
    parts = date_str.split("-")

    # 解析 年/月/日
    if parts[0] == "":  # 处理公元前：-000156-01-01
        year = -int(parts[1])
        month = int(parts[2])
        day = int(parts[3])
    else:
        year = int(parts[0])
        month = int(parts[1])
        day = int(parts[2])

    # 天文算法：计算儒略日 JDN（0点修正版）
    a = (14 - month) // 12
    y = year + 4800 - a
    m = month + 12 * a - 3

    jdn = day + (153 * m + 2) // 5 + 365 * y + y // 4 - y // 100 + y // 400 - 32045
    jdn = jdn - 0.5  # 🔥 关键修复！从正午 → 0点
    
    # 转 UTC Unix 秒
    timestamp = (jdn - 2440587.5) * 86400
    return int(round(timestamp))

with open("chaodai.json", "r", encoding="utf-8") as f:
    data = json.load(f)
    for x in data:
        id = x["id"] if "id" in x else None
        start = date_to_utc_timestamp(x["start"])  if "start" in x else None
        end =  date_to_utc_timestamp(x["end"])  if "end" in x else None
        content = x["content"]  if "content" in x else None
        title = x["title"]  if "title" in x else None
        group = 1
        subgroup = 0
        type = "range"
        color1= x["color1"]  if "color1" in x else None
        color2 = x["color2"]  if "color2" in x else None
        precision_start = x["precision_start"]  if "precision_start" in x else None
        precision_end = x["precision_end"]  if "precision_end" in x else None
        conn.execute("insert into item(start, end, content, title, group_id, subgroup_id, type, color1, color2,precision_start,precision_end) values (?, ?, ?, ?, ?, ?, ?, ?, ?,?,?)", (start, end, content, title, group, subgroup, type, color1, color2,precision_start,precision_end))
    conn.commit()