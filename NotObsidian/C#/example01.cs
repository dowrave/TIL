class Computer
{
}

interface IMonitor
{
	void TurnOn();
}

interface IKeyboard {}

class Notebook : Computer, IMonitor, IKeyboard
{
	public void TurnOn() {} // 추상 메서드와 달리 override를 쓰지 않는다.
}