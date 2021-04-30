# TAML

Because Tabs are just better.

## Why?

> Tabs have been outlawed since they are treated differently by different editors and tools. And since indentation is so critical to proper interpretation of YAML, this issue is just too tricky to even attempt. Indeed Guido van Rossum of Python has acknowledged that allowing TABs in Python source is a headache for many people and that were he to design Python again, he would forbid them.

_[YAML FAQ](https://yaml.org/faq.html)_

> Tabs are treated differently by different editors and tools

So what?

> And since indentation is so critical to proper interpretation of YAML

Exactly, a Tab is always a single level of indentation, but YAML gives you the choice 2 OR 4 spaces! (uniformed through the document)

### Personal Reason

I use Tabs in everything. When I use YAML, VSCode doesn't Switch to spaces, so here's that.

## What?

TAML is YAML, but Space indentations are prohibited, since they're the second source of evil. Instead, Tabs are Used.

## Compatability?

It uses [go-yaml](https://github.com/go-yaml/yaml) under the hood, so it all depends on that package,which currently supports YAML 1.2.
