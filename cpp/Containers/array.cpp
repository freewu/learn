#include <iostream>
#include <array>

int main() {	
	std::array<int ,5> arr = {2,3,4,5,6};
	// normal
	std::cout << "arr contains:";
	for (auto i = arr.begin(); i != arr.end(); ++i) {
		std::cout << ' ' << *i;	
	}
	std::cout << '\n';
	
	// at
	std::cout << "arr contains:";
	for (int i = 0; i < 5; i++) {
		std::cout << ' ' << arr.at(i);	
	}
	std::cout << '\n';

	// const iterator
	std::cout << "arr contains:";
	for (auto i = arr.cbegin(); i != arr.cend(); ++i) {
		std::cout << ' ' << *i;	 // cannot modify *i
	}
	std::cout << '\n';

	// reverse iterator
	std::cout << "arr contains:";
	for (auto i = arr.rbegin(); i != arr.rend(); ++i) {
		std::cout << ' ' << *i;	
	}
	std::cout << '\n';

	// reverse const iterator
	std::cout << "arr contains:";
	for (auto i = arr.crbegin(); i != arr.crend(); ++i) {
		std::cout << ' ' << *i;	 // cannot modify *i
	}
	std::cout << '\n';
	
	// front back
	std::cout << "front is:" << arr.front() << '\n';
	std::cout << "back is:" << arr.back() << '\n';

	arr.back() = 22;
	arr.front() = 33;
	std::cout << "arr contains:";
	for (auto i = arr.begin(); i != arr.end(); ++i) {
		std::cout << ' ' << *i;	
	}
	std::cout << '\n';
	
	
	
	return 0;
}
