// 10보다 작은 3과 5의 배수를 모두 더하면 23일 때, 1000보다 작은 자연수 중 3 또는 5의 배수를 모두 더하면 얼마일까요?

using System;
using learn;

namespace learn
    {
        class Program{
        static void Main(string[] args) 
        {
            int result = 0;
            int target = 10;

            for (int i = 1; i < target; i ++)
            {
                if (i % 3 == 0 || i % 5 == 0)
                {
                    result += i;
                }
            }
            Console.WriteLine(result);

        }
    }
    }