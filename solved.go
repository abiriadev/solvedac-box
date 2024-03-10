package main

type User struct {
  handle string
  bio string
  solvedCount int
  voteCount int
  tier int
  rating int
  class int
  classDecoration string
  rivalCount int
  reverseRivalCount int
  rank int
}

func fetchUserData(user string) (User, error) {}
  res, err := http.Get("https://solved.ac/api/v3/user/show?handle=" + user)
  if err != nil {
    return User{}, err
  }

  return User{}, nil
}
