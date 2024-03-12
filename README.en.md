<p align="center"><img src="./assets/box.png" alt="gist box picture" width="500"></p>

<h1 align="center">Solved.ac Box</h1>

<p align="center"><img alt="GitHub forks" src="https://img.shields.io/github/forks/abiriadev/solvedac-box?style=flat-square&logo=github&color=17ce3a"> <img alt="GitHub Actions Workflow Status" src="https://img.shields.io/github/actions/workflow/status/abiriadev/solvedac-box/update.yaml?style=flat-square&logo=github&label=Update&color=17ce3a">
 <img alt="GitHub Gist last commit" src="https://img.shields.io/github/gist/last-commit/7f2692a91fbddcafd096d52d28f799c5?style=flat-square&logo=github&label=Gist%20updated&color=17ce3a">

<p align="center">
  📊 Show your <a href="https://solved.ac">Solved.ac</a> profile information as a GitHub Gist 📊
</p>

---

## 🎒 Prep Work

1. [Create a new GitHub PAT token](https://github.com/settings/personal-access-tokens/new) with `Gists` to `Read and write` and copy it.
2. Create a new public [GitHub Gist](https://gist.github.com/) then copy the ID from the URL.
3. Lastly, prepare your BOJ(Baekjoon Online Judge) account ID.

## 🖥 Project Setup

1. [Fork this repository](https://github.com/abiriadev/solvedac-box/fork).
2. Go to `Settings` > `Secrets and variables` > `Actions` then add the below informations as a `New repository secret`s.
3. Go to `Actions` > `Update Gist` then push the `Run workflow` > `Run workflow` button!
4. After the Gist has beed updated, you can pin it to your GitHub profile!

## 🤫 Environments Secrets

-   **GH_TOKEN:** The GitHub PAT token generated above
-   **GIST_ID:** The ID of your GitHub Gist
-   **USERNAME:** Your BOJ account ID

## 📄 License

[![GitHub](https://img.shields.io/github/license/abiriadev/pia?color=17ce3a&style=for-the-badge)](./LICENSE)

_<sub>Special thanks to [BOJ](https://www.acmicpc.net/) and [@solved-ac](https://github.com/solved-ac)</sub>_
