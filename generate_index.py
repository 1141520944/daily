#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
自动生成项目索引到README.md（包括算法题解和文档）
"""

import os
import re
from pathlib import Path


def get_algorithm_files(leetcode_dir):
    """
    获取leetcode目录下所有算法文件（排除base.go等基础文件）
    返回按文件名排序的文件列表
    """
    algorithm_files = []
    
    if not os.path.exists(leetcode_dir):
        return algorithm_files
    
    for file in os.listdir(leetcode_dir):
        file_path = os.path.join(leetcode_dir, file)
        
        # 只处理.go文件，排除base.go
        if file.endswith('.go') and file != 'base.go':
            algorithm_files.append(file)
    
    # 按文件名排序
    algorithm_files.sort()
    return algorithm_files


def get_document_files(doc_dir):
    """
    获取doc目录下所有Markdown文档
    返回按路径排序的文件列表
    """
    document_files = []
    
    if not os.path.exists(doc_dir):
        return document_files
    
    # 递归遍历doc目录
    for root, dirs, files in os.walk(doc_dir):
        for file in files:
            if file.endswith('.md'):
                # 获取相对路径
                full_path = os.path.join(root, file)
                rel_path = os.path.relpath(full_path, doc_dir)
                document_files.append(rel_path)
    
    # 按路径排序
    document_files.sort()
    return document_files


def extract_doc_info(rel_path):
    """
    从文档路径中提取信息
    例如: "算法/回溯算法模板总结.md" -> {"category": "算法", "title": "回溯算法模板总结"}
    """
    # 统一路径分隔符为正斜杠
    rel_path = rel_path.replace('\\', '/')
    
    # 移除.md后缀
    name_without_ext = rel_path.replace('.md', '')
    
    # 分割路径
    parts = name_without_ext.split('/')
    
    if len(parts) > 1:
        return {
            'category': parts[-2],  # 倒数第二个是分类
            'title': parts[-1],     # 最后一个是文件名
            'path': rel_path
        }
    else:
        return {
            'category': '其他',
            'title': parts[0],
            'path': rel_path
        }


def generate_readme_content(algorithm_files, document_files, leetcode_dir, doc_dir):
    """
    生成完整的README内容
    """
    content = "# 📚 学习记录与算法题解\n\n"
    
    # 第一部分：算法题解
    content += "## 💻 LeetCode 算法题解\n\n"
    
    if algorithm_files:
        content += "| 题号 | 题目名称 | 链接 |\n"
        content += "|------|----------|------|\n"
        
        for filename in algorithm_files:
            # 提取题号和标题
            name_without_ext = filename.replace('.go', '')
            match = re.match(r'^(\d+)\.\s*(.+)$', name_without_ext)
            
            if match:
                number = match.group(1)
                title = match.group(2).strip()
                content += f"| {number} | {title} | [查看代码](leetcode/{filename}) |\n"
            else:
                content += f"| - | {name_without_ext} | [查看代码](leetcode/{filename}) |\n"
        
        content += f"\n**算法题总计**: {len(algorithm_files)} 道\n"
    else:
        content += "暂无算法题解文件。\n"
    
    # 第二部分：技术文档
    content += "\n---\n\n"
    content += "## 📖 技术文档\n\n"
    
    if document_files:
        # 按分类组织
        categories = {}
        for rel_path in document_files:
            info = extract_doc_info(rel_path)
            category = info['category']
            if category not in categories:
                categories[category] = []
            categories[category].append(info)
        
        # 输出每个分类
        for category in sorted(categories.keys()):
            content += f"### {category}\n\n"
            content += "| 文档名称 | 链接 |\n"
            content += "|----------|------|\n"
            
            for doc in categories[category]:
                # 计算相对路径（从项目根目录开始）
                link_path = f"doc/{doc['path']}".replace('\\', '/')
                content += f"| {doc['title']} | [查看文档]({link_path}) |\n"
            
            content += "\n"
        
        content += f"**文档总计**: {len(document_files)} 篇\n"
    else:
        content += "暂无技术文档。\n"
    
    return content


def update_readme(readme_path, leetcode_dir, doc_dir):
    """
    更新README.md文件
    """
    # 获取文件列表
    algorithm_files = get_algorithm_files(leetcode_dir)
    document_files = get_document_files(doc_dir)
    
    # 生成完整内容
    content = generate_readme_content(algorithm_files, document_files, leetcode_dir, doc_dir)
    
    # 写入README.md
    with open(readme_path, 'w', encoding='utf-8') as f:
        f.write(content)
    
    total_items = len(algorithm_files) + len(document_files)
    print(f"✓ 已更新 README.md")
    print(f"  - 算法题解: {len(algorithm_files)} 道")
    print(f"  - 技术文档: {len(document_files)} 篇")
    print(f"  - 总计: {total_items} 项")
    
    # 打印生成的内容预览
    print("\n生成的索引内容：")
    print("=" * 60)
    print(content)
    print("=" * 60)


def main():
    """
    主函数
    """
    # 获取脚本所在目录
    script_dir = Path(__file__).parent.absolute()
    
    # 设置路径
    readme_path = os.path.join(script_dir, 'README.md')
    leetcode_dir = os.path.join(script_dir, 'leetcode')
    doc_dir = os.path.join(script_dir, 'doc')
    
    # 检查目录是否存在
    if not os.path.exists(leetcode_dir) and not os.path.exists(doc_dir):
        print(f"错误: leetcode和doc目录都不存在")
        return
    
    if not os.path.exists(readme_path):
        print(f"创建新的 README.md 文件")
    
    # 更新README
    update_readme(readme_path, leetcode_dir, doc_dir)


if __name__ == '__main__':
    main()
