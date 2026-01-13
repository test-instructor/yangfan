import random
import string
import json


# 生成随机字符串（原函数保留）
def generate_random_string(length):
    """生成指定长度的随机字符串（字母+数字）"""
    letters_and_digits = string.ascii_letters + string.digits
    return ''.join(random.choice(letters_and_digits) for _ in range(length))


# 生成随机中文姓名
def generate_chinese_name():
    """生成随机中文姓名"""
    # 常见姓氏
    first_names = ['王', '李', '张', '刘', '陈', '杨', '黄', '赵', '吴', '周',
                   '徐', '孙', '马', '朱', '胡', '林', '郭', '何', '高', '罗']
    # 常见名字
    last_names = ['伟', '芳', '娜', '敏', '静', '强', '磊', '洋', '杰', '娟',
                  '涛', '明', '华', '丽', '红', '军', '平', '刚', '勇', '艳']

    # 随机生成2-3个字的名字
    if random.random() > 0.3:
        # 2个字的名字
        return random.choice(first_names) + random.choice(last_names)
    else:
        # 3个字的名字
        return random.choice(first_names) + random.choice(last_names) + random.choice(last_names)


# 生成随机邮箱
def generate_email(name, phone):
    """根据姓名和手机号生成真实感的邮箱"""
    # 常见邮箱域名
    domains = ['@qq.com', '@163.com', '@126.com', '@gmail.com', '@sina.com', '@sohu.com']
    # 姓名拼音（简化处理）
    name_pinyin_map = {
        '王': 'wang', '李': 'li', '张': 'zhang', '刘': 'liu', '陈': 'chen',
        '杨': 'yang', '黄': 'huang', '赵': 'zhao', '吴': 'wu', '周': 'zhou',
        '伟': 'wei', '芳': 'fang', '娜': 'na', '敏': 'min', '静': 'jing',
        '强': 'qiang', '磊': 'lei', '洋': 'yang', '杰': 'jie', '娟': 'juan'
    }

    # 提取姓名拼音（如果不在映射表则用随机字母）
    name_pinyin = ''
    for char in name:
        name_pinyin += name_pinyin_map.get(char, random.choice(['a', 'b', 'c', 'd', 'e', 'f', 'g']))

    # 邮箱前缀组合方式
    prefix_options = [
        f"{name_pinyin}{random.randint(1980, 2000)}",  # 姓名+出生年份
        f"{phone[:6]}{random.choice(['', '_', ''])}",  # 手机号前6位
        f"{name_pinyin[:3]}{random.randint(100, 999)}"  # 姓名前3字母+随机数
    ]

    return random.choice(prefix_options) + random.choice(domains)


# 生成随机地址
def generate_address():
    """生成随机中文地址"""
    provinces = ['北京市', '上海市', '广东省', '江苏省', '浙江省', '四川省', '山东省', '河南省']
    cities = ['朝阳区', '浦东新区', '深圳市', '苏州市', '杭州市', '成都市', '青岛市', '郑州市']
    streets = ['街道', '路', '大道', '巷', '胡同']

    province = random.choice(provinces)
    city = random.choice(cities)
    street = f"{random.randint(1, 999)}{random.choice(streets)}"
    community = f"{random.choice(['阳光', '幸福', '锦绣', '花园', '新城', '世纪'])}小区{random.randint(1, 50)}栋{random.randint(1, 30)}单元{random.randint(101, 999)}室"

    return f"{province}{city}{street}{community}"


# 生成随机钱包余额
def generate_balance():
    """生成随机钱包余额（0.01 - 99999.99）"""
    return round(random.uniform(0.01, 99999.99), 2)


# 主函数
def create_data(ENV, last_data, number):
    last_data_old = last_data
    phone_number = int(last_data.get("phone", 0))  # 增加默认值，避免KeyError
    if phone_number == 0:
        phone_number = 1990000000

    l = []
    for k in range(number):
        # 基础信息
        current_phone = phone_number + k + 1
        name = generate_chinese_name()

        # 构建用户数据
        user_data = {
            "phone": current_phone,
            "password": generate_random_string(random.randint(10, 18)),
            "ENV_name": ENV.get("ENV_name"),
            "number": k,
            # 新增用户属性
            "name": name,
            "gender": random.choice(['男', '女']),  # 随机性别
            "email": generate_email(name, str(current_phone)),  # 基于姓名和手机号的邮箱
            "age": random.randint(18, 60),  # 随机年龄
            "balance": generate_balance(),  # 钱包余额
            "address": generate_address(),  # 详细地址
            "id_card": f"{random.randint(100000, 999999)}{random.randint(1950, 2005)}{random.randint(10, 12)}{random.randint(10, 31)}{random.randint(100, 999)}{random.choice([0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 'X'])}",
            # 模拟身份证号
            "register_time": f"20{random.randint(10, 24)}-{random.randint(1, 12):02d}-{random.randint(1, 28):02d} {random.randint(0, 23):02d}:{random.randint(0, 59):02d}:{random.randint(0, 59):02d}",
            # 注册时间
            "user_level": random.choice(['普通会员', 'VIP会员', '钻石会员']),  # 用户等级
            "is_vip": random.choice([True, False]),  # 是否VIP
            "login_count": random.randint(1, 500)  # 登录次数
        }
        l.append(user_data)

    # 最后一条数据作为last_data
    data = {"list": l, "last_data": l[-1] if l else last_data_old}

    return data


def update_data(ENV,number,data_old):
    data = [{},{}]
    return  data