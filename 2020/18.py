import re
from dataclasses import dataclass
from typing import Optional, Union

inputs = [line.replace(' ', '').strip() for line in open('input18.txt') if line]

@dataclass
class Token:
    type: str
    val: Optional[Union[int, str]] = None
    prec: Optional[int] = None

class peekable:
    def __init__(self, it):
        self.it = iter(it)
        self.cache = []

    def peek(self):
        if not self.cache:
            try:
                self.cache.append(next(self.it))
            except StopIteration:
                return None
        return self.cache[-1]

    def __next__(self):
        if self.cache:
            return self.cache.pop()
        return next(self.it)

    def __iter__(self):
        return self

def tokens(input, precedence):
    input = input.strip().replace(' ', '')
    i = 0
    while i < len(input):
        c = input[i]
        if c in '+*':
            yield Token('op', c, precedence[c])
            i += 1
        elif c in '(':
            yield Token('lparen')
            i += 1
        elif c in ')':
            yield Token('rparen')
            i += 1
        elif m := re.match(r'\d+', input[i:]):
            yield Token('num', int(m.group()))
            i += m.end()
        else:
            raise Exception(f'unexpected token {input[i:]!r}')

def atom(tok):
    t = next(tok)
    if t.type == 'num':
        return t.val

    assert t.type == 'lparen', f'unexpected token {t!r}'
    v = expr(tok, 1)
    assert next(tok).type == 'rparen'
    return v

def expr(tok, min_prec):
    r = atom(tok)
    while (t := tok.peek()) and t.type == 'op' and t.prec >= min_prec:
        next(tok)
        rhs = expr(tok, t.prec + 1) # +1 means left-assoc
        assert t.val in '+*', f'unexpected op {t.val!r}'
        r = (r * rhs) if t.val == '*' else (r + rhs)
    return r

for precedence in [{'*':1,'+':1}, {'*':1,'+':2}, {'*':2,'+':1}]:
    print(sum(expr(peekable(tokens(I, precedence)), 1) for I in inputs))
