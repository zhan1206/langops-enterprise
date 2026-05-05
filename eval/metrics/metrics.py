"""LangOps Enterprise - Evaluation Metrics Library"""

from dataclasses import dataclass
from typing import Optional


@dataclass
class MetricResult:
    name: str
    score: float
    details: str = ""


def accuracy_score(output: str, expected: str) -> MetricResult:
    """Compute accuracy between output and expected."""
    if not output or not expected:
        return MetricResult("accuracy", 0.0, "Empty input")
    correct = sum(1 for a, b in zip(output, expected) if a == b)
    total = max(len(output), len(expected))
    score = correct / total if total > 0 else 0.0
    return MetricResult("accuracy", round(score, 4))


def relevance_score(output: str, query: str) -> MetricResult:
    """Compute relevance of output to query."""
    if not output or not query:
        return MetricResult("relevance", 0.0, "Empty input")
    query_terms = set(query.lower().split())
    output_terms = set(output.lower().split())
    overlap = len(query_terms & output_terms)
    score = overlap / len(query_terms) if query_terms else 0.0
    return MetricResult("relevance", round(min(score, 1.0), 4))


def factuality_score(output: str, context: str) -> MetricResult:
    """Compute factual consistency with context."""
    if not output or not context:
        return MetricResult("factuality", 0.0, "Empty input")
    output_terms = set(output.lower().split())
    context_terms = set(context.lower().split())
    supported = len(output_terms & context_terms)
    score = supported / len(output_terms) if output_terms else 0.0
    return MetricResult("factuality", round(min(score, 1.0), 4))


def fluency_score(output: str) -> MetricResult:
    """Compute fluency of output text."""
    if not output:
        return MetricResult("fluency", 0.0, "Empty output")
    words = output.split()
    avg_len = sum(len(w) for w in words) / len(words) if words else 0
    score = min(avg_len / 6.0, 1.0)
    return MetricResult("fluency", round(score, 4))


def harmfulness_score(output: str) -> MetricResult:
    """Detect harmful content in output."""
    harmful_patterns = ["violence", "illegal", "harmful", "dangerous"]
    output_lower = output.lower()
    found = sum(1 for p in harmful_patterns if p in output_lower)
    score = 1.0 - (found / len(harmful_patterns)) if harmful_patterns else 1.0
    return MetricResult("harmfulness", round(score, 4))