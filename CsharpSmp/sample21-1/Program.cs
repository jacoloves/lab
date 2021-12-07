using System;

public interface IEnumerator
{
    IEnumerator GetEnumerator();
    object Current { get; }
    bool MoveNext();
    void Reset();
}

public interface IEnumerator<out T> : IDisposable, IEnumerator
{
    IEnumerator<T> GetEnumerator();
    T Current { get; }
}
public class List<T> : IEnumerable<T>
{
    T[] _array;
    int _count;

    public IEnumerator<T> GetEnumerator()
    {
        return new ListEnumerator(this);
    }

    class ListEnumerator : IEnumerator<T>
    {
        List<T> _list;
        int _i;

        public ListEnumerator(List<T> list)
        {
            _list = list;
            _i = -1;
        }

        public bool MoveNext()
        {
            ++_i;
            return _i < _list._count;
        }

        public T Current
        {
            get
            {
                return _list._array[_i];
            }
        }
    }

    class AbcEnumerator : IEnumerator<string>
    {
        int _i = 0;
        public bool MoveNext()
        {
            switch (_i)
            {
                case 0:
                    ++_i;
                    Current = "a";
                    return true;
                case 1:
                    ++_i;
                    Current = "b";
                    return true;
                case 2:
                    ++_i;
                    Current = "c";
                    return true;
                default:
                    return false;
            }
        }
    }

    public string Current { get; private set; }
}