import seaborn as sns
# import pandas as pd
import matplotlib.pyplot as plt

df = sns.load_dataset("iris")

sns.pairplot(df, hue="species", diag_kind="kde")
plt.show()
