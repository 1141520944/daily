#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
自动生成LeetCode算法题解索引到README.md
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
    
    for file in os.listdir(leetcode_dir):
        file_path = os.path.join(leetcode_dir, file)
        
        # 只处理.go文件，排除base.go
        if file.endswith('.go') and file != 'base.go':
            algorithm_files.append(file)
    
    # 按文件名排序
    algorithm_files.sort()
    return algorithm_files


def extract_problem_info(filename):
    """
    从文件名中提取题目信息
    例如: "104. 二叉树的最大深度.go" -> {"number": "104", "title": "二叉树的最大深度"}
    """
    # 移除.go后缀
    name_without_ext = filename.replace('.go', '')
    
    # 匹配格式: "数字. 题目名称"
    match = re.match(r'^(\d+)\.\s*(.+)$', name_without_ext)
    
    if match:
        return {
            'number': match.group(1),
            'title': match.group(2).strip(),
            'filename': filename
        }
    else:
        # 如果格式不匹配，返回原始文件名
        return {
            'number': '',
            'title': name_without_ext,
            'filename': filename
        }


def generate_markdown_table(algorithm_files, leetcode_dir):
    """
    生成Markdown格式的表格
    """
    if not algorithm_files:
        return "# LeetCode 算法题解\n\n暂无算法题解文件。\n"
    
    # 表头
    markdown = "# LeetCode 算法题解\n\n"
    markdown += "| 题号 | 题目名称 | 链接 |\n"
    markdown += "|------|----------|------|\n"
    
    # 表格内容
    for filename in algorithm_files:
        info = extract_problem_info(filename)
        
        # 生成相对路径链接
        link_path = f"leetcode/{filename}"
        
        if info['number']:
            markdown += f"| {info['number']} | {info['title']} | [查看代码]({link_path}) |\n"
        else:
            markdown += f"| - | {info['title']} | [查看代码]({link_path}) |\n"
    
    return markdown


def update_readme(readme_path, leetcode_dir):
    """
    更新README.md文件
    """
    # 获取算法文件列表
    algorithm_files = get_algorithm_files(leetcode_dir)
    
    # 生成Markdown内容
    content = generate_markdown_table(algorithm_files, leetcode_dir)
    
    # 添加统计信息
    total = len(algorithm_files)
    content += f"\n\n**总计**: {total} 道题\n"
    
    # 写入README.md
    with open(readme_path, 'w', encoding='utf-8') as f:
        f.write(content)
    
    print(f"✓ 已更新 README.md，共 {total} 道算法题")
    
    # 打印生成的内容预览
    print("\n生成的索引内容：")
    print("-" * 50)
    print(content)
    print("-" * 50)


def main():
    """
    主函数
    """
    # 获取脚本所在目录
    script_dir = Path(__file__).parent.absolute()
    
    # 设置路径
    readme_path = os.path.join(script_dir, 'README.md')
    leetcode_dir = os.path.join(script_dir, 'leetcode')
    
    # 检查目录是否存在
    if not os.path.exists(leetcode_dir):
        print(f"错误: leetcode目录不存在: {leetcode_dir}")
        return
    
    if not os.path.exists(readme_path):
        print(f"创建新的 README.md 文件")
    
    # 更新README
    update_readme(readme_path, leetcode_dir)


if __name__ == '__main__':
    main()
