package main

func fetchUserData(user string) (User, error) {}
  res, err := http.Get("https://solved.ac/api/v3/user/show?handle=" + user)
  if err != nil {
  return User{}, err
  }
}
