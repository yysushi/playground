# Notes for printf and cout

## from the internet

- printf vs cout

正直に言えば、printf も cout も決して最新の C++ を代表するものではありません。printf 関数は可変個の引数を受け取る関数の一例で、C プログラミング言語から引き継がれたやや脆弱な機能の数少ない有効な使い方の 1 つです。可変個の引数を受け取る関数は、可変個引数テンプレートより前から存在します。可変個引数テンプレートは、可変個の型や引数を処理するために、まさに最新かつ堅牢な機能を提供しています。これとは対照的に、cout は可変個の引数を使用しませんが、代わりに、仮想関数の呼び出しに大きく依存します。コンパイラは仮想関数呼び出しのパフォーマンスを必ずしも最適化できるわけではありません。実際、CPU 設計の進化は cout のポリモーフィックなアプローチのパフォーマンスを向上することはほとんどなく、printf に有利に働いています。したがって、パフォーマンスと効率性を求めるのであれば、printf を選択するのが賢明です。また、printf の方が生成されるコードが簡潔になります。

<https://docs.microsoft.com/ja-jp/archive/msdn-magazine/2015/march/windows-with-c-using-printf-with-modern-c>

- iostream and stdio.h

  - stio.h is the header file in the C standard library, which is used for input/output.
  - iostream is the input output class in C++.

<https://stackoverflow.com/questions/28764438/what-the-difference-between-stdio-h-and-iostream?answertab=votes#tab-top>
