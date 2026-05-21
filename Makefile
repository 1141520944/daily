.PHONY: help index clean

# 默认目标
help:
	@echo "可用命令:"
	@echo "  make index   - 生成算法题解索引到 README.md"
	@echo "  make clean   - 清理生成的索引文件"
	@echo "  make help    - 显示帮助信息"

# 生成算法题解索引
index:
	@python generate_index.py

# 清理索引（可选）
clean:
	@echo "清理 README.md..."
	@echo "" > README.md
	@echo "✓ 已清理 README.md"
