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
cipherCode  -help
```

Afin de voir les commandes disponibles 

### Exemple :

#### Afin de coder le message voulu:
```
cipherCode -code
```

![cipherCode-code](https://github.com/user-attachments/assets/1a684f71-6d0f-402e-9f1b-9646299e532a)


#### Afin de décoder le message voulu:
```
cipherCode -decode
```


![cipherCode-decode](https://github.com/user-attachments/assets/8f2d80bf-5d40-4cd4-91b8-12126bc7edd4)

#### Afin de coder le message voulu avec le delta choisi (Nombre de décalage de lettres):
```
cipherCode -code -delta 1
```


![cipherCode-code-delta-1](https://github.com/user-attachments/assets/955e80d0-050e-47e2-865e-75e29b4a491b)


#### Afin de décoder le message voulu avec le delta choisi (Nombre de décalage de lettres):
```
cipherCode -decode -delta 1
```

![cipherCode-decode-delta-1](https://github.com/user-attachments/assets/ab2d7723-d427-42b2-9252-ae4d7caf927f)

