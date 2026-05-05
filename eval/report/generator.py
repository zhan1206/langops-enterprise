"""LangOps Enterprise - Evaluation Report Generator"""

import json
from datetime import datetime
from typing import Any


class ReportGenerator:
    """Generate evaluation reports in multiple formats."""

    def generate_json(self, eval_report: dict) -> str:
        """Generate JSON format report."""
        return json.dumps(eval_report, indent=2, ensure_ascii=False)

    def generate_markdown(self, eval_report: dict) -> str:
        """Generate Markdown format report."""
        lines = [
            f"# Evaluation Report",
            f"",
            f"**ID**: {eval_report.get('id', 'N/A')}",
            f"**Date**: {datetime.now().isoformat()}",
            f"**Total Score**: {eval_report.get('total_score', 0)}",
            f"**Pass Rate**: {eval_report.get('pass_rate', 0) * 100:.1f}%",
            f"",
            f"## Dimension Results",
            f"",
            f"| Dimension | Score | Weight | Passed |",
            f"|-----------|-------|--------|--------|",
        ]
        for r in eval_report.get("results", []):
            passed_mark = "✅" if r.get("passed") else "❌"
            lines.append(
                f"| {r.get('dimension', 'N/A')} | {r.get('score', 0):.2f} | "
                f"{r.get('weight', 0):.2f} | {passed_mark} |"
            )
        return "\n".join(lines)

    def generate_html(self, eval_report: dict) -> str:
        """Generate HTML format report."""
        return f"<html><body><h1>Eval Report - {eval_report.get('id', 'N/A')}</h1></body></html>"