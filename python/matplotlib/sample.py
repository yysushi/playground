import numpy as np
import matplotlib.pyplot as plt

x = np.linspace(0, 2, 100)


plt.plot(x, x, label='linear')
plt.plot(x, x**2, label='quadratic')
plt.plot(x, x**3, label='cubic')

plt.plot([0.25, 1.0, 1.5], [0.5, 2.0, 4.0], label='plot')

plt.xlabel('x label')
plt.ylabel('y label')

plt.title("Simple Plot")

plt.legend()

plt.show()
