#!/usr/bin/env python3
"""Regenerate the progress section of README.md from solutions/go/.

Exercise names, concepts and difficulty come from the Exercism Go track's
config.json. Only the text between the PROGRESS markers is rewritten.

Usage:
  python scripts/update_readme.py                 # fetch config.json from GitHub
  python scripts/update_readme.py --config PATH   # use a local copy (offline)
"""
import argparse
import json
import os
import re
import subprocess
import sys
import urllib.request
from datetime import date

CONFIG_URL = "https://raw.githubusercontent.com/exercism/go/main/config.json"
SOLUTIONS = os.path.join("solutions", "go")
README = "README.md"
START = "<!-- PROGRESS:START -->"
END = "<!-- PROGRESS:END -->"


def load_config(path):
    if path:
        with open(path) as f:
            return json.load(f)
    with urllib.request.urlopen(CONFIG_URL, timeout=30) as r:
        return json.load(r)


def last_updated():
    try:
        out = subprocess.run(
            ["git", "log", "-1", "--format=%cs", "--", SOLUTIONS],
            capture_output=True, text=True, check=True,
        ).stdout.strip()
        if out:
            return out
    except (subprocess.CalledProcessError, FileNotFoundError):
        pass
    return date.today().isoformat()


def latest_iteration(slug):
    its = [d for d in os.listdir(os.path.join(SOLUTIONS, slug)) if d.isdigit()]
    return max(its, key=int) if its else ""


def link(slug, name):
    it = latest_iteration(slug)
    path = f"{SOLUTIONS}/{slug}/{it}/" if it else f"{SOLUTIONS}/{slug}/"
    return f"[{name}]({path})"


def tags(xs):
    return ", ".join(f"`{x}`" for x in xs) or "—"


def build_section(config):
    done = sorted(
        d for d in os.listdir(SOLUTIONS)
        if os.path.isdir(os.path.join(SOLUTIONS, d))
    )
    concept = config["exercises"]["concept"]
    practice = {e["slug"]: e for e in config["exercises"]["practice"]}
    by_concept = {e["slug"]: e for e in concept}

    concept_done = [e["slug"] for e in concept if e["slug"] in done]
    practice_done = sorted(
        (s for s in done if s in practice),
        key=lambda s: (practice[s].get("difficulty", 0), practice[s]["name"]),
    )
    unknown = [s for s in done if s not in by_concept and s not in practice]

    L = []
    L.append(
        f"**Progress:** {len(done)} exercises completed "
        f"({len(concept_done)} concept · {len(practice_done)} practice) — "
        f"last updated {last_updated()}."
    )
    L += ["", "## Concept exercises", "",
          "Guided exercises that each introduce one Go language feature, in track order.",
          "", "| Exercise | Concepts |", "|---|---|"]
    for s in concept_done:
        e = by_concept[s]
        L.append(f"| {link(s, e['name'])} | {tags(e.get('concepts', []))} |")
    L += ["", "## Practice exercises", "",
          "Open-ended problems that apply those concepts, sorted by Exercism's difficulty rating (1–10).",
          "", "| Exercise | Difficulty | Practices |", "|---|---|---|"]
    for s in practice_done:
        e = practice[s]
        L.append(
            f"| {link(s, e['name'])} | {e.get('difficulty', '—')} | "
            f"{tags(e.get('practices', [])[:3])} |"
        )
    if unknown:
        L += ["", "## Other", ""] + [f"- {link(s, s)}" for s in unknown]
    return "\n".join(L)


def main():
    ap = argparse.ArgumentParser()
    ap.add_argument("--config", help="local path to the Go track config.json")
    args = ap.parse_args()

    with open(README) as f:
        readme = f.read()
    if START not in readme or END not in readme:
        sys.exit(f"README.md is missing {START} / {END} markers")

    section = build_section(load_config(args.config))
    new = re.sub(
        re.escape(START) + r".*?" + re.escape(END),
        lambda _: f"{START}\n{section}\n{END}",
        readme, flags=re.S,
    )
    if new != readme:
        with open(README, "w") as f:
            f.write(new)
        print("README.md updated")
    else:
        print("README.md already up to date")


if __name__ == "__main__":
    main()
