
# Go Web Server Project

Ce projet est un serveur web simple écrit en Go (Golang). Il comprend des routes de base pour la page d'accueil et une page "À propos".

<img src="images.png" />
## Structure du projet

```txt
GoWeb/
├── main.go
├── handlers.go
├── handlers_test.go
```

- **main.go**: Le fichier principal qui démarre le serveur.
- **handlers.go**: Contient les handlers pour les différentes routes.
- *)


## Exécution du serveur

Pour compiler et exécuter le serveur:

```bash
go run main.go
```

Le serveur démarre sur le port `8080`. Vous pouvez maintenant accéder aux routes suivantes dans votre navigateur ou en utilisant `curl`:

- `http://localhost:8080/` : Accède à la page d'accueil.
- `http://localhost:8080/about` : Accède à la page "À propos".

