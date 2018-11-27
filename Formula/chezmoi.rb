class Chezmoi < Formula
  desc "chezmoi is a tool for managing your dotfiles across multiple machines."
  homepage "https://github.com/twpayne/chezmoi"
  url "https://github.com/joemiller/chezmoi/releases/download/v0.0.2/chezmoi_0.0.2_darwin_amd64.tar.gz"
  version "0.0.2"
  sha256 "281bc0f97a0fea527bb56b469887de42848d9a35dc1bacd1e21ce5329f5bbe4f"

  def install
    bin.install "chezmoi"
  end
end
