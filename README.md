# CipherCode

Le chiffre de César est un type de chiffrement par substitution simple et largement connu, où chaque lettre de l’alphabet est remplacée par une lettre à une position fixe plus basse dans l’alphabet. En français, il s’applique uniquement aux lettres de l’alphabet latin.

## Principe

Le chiffre de César consiste à décaler chaque lettre de l’alphabet de plusieurs positions avant de l’écrire. Par exemple, si le décalage est de 3, la lettre “A” devient “D”, “B” devient “E”, et ainsi de suite.


### Installation avec Homebrew :

Tapez :

```
brew tap Lachignol/homebrew-Lachignol 
```

```
brew install lachignol/lachignol/cipherCode  
```

Vous pouvez ensuite taper :

```
cipherCode  --help
```

Afin de voir les commandes disponibles 

### Exemples :

#### Afin de coder le message voulu:
```
cipherCode --code
```

![cipherCode-code](https://github.com/user-attachments/assets/7f05c8ff-6b6d-4976-904f-4fc464b5ffbf)


#### Afin de décoder le message voulu:
```
cipherCode --decode
```

![cipherCode-decode](https://github.com/user-attachments/assets/ec3f87f6-2ef9-4076-92eb-a26eaa899913)


#### Afin de coder le message voulu avec le delta choisi (Nombre de décalage de lettres):
```
cipherCode --code --delta 1
```

![cipherCode-code-delta-1](https://github.com/user-attachments/assets/56ffc143-9485-45b5-a014-e472de1a6ac9)


#### Afin de décoder le message voulu avec le delta choisi (Nombre de décalage de lettres):
```
cipherCode --decode --delta 1
```

![cipherCode-decode-delta-1](https://github.com/user-attachments/assets/523f0493-da1a-41fb-a6f4-6f62123768c0)


