class Chezmoi < Formula
  desc "chezmoi is a tool for managing your dotfiles across multiple machines."
  homepage "https://github.com/twpayne/chezmoi"
  url "https://github.com/joemiller/chezmoi/releases/download/v0.0.2/chezmoi_0.0.2_darwin_amd64.tar.gz"
  version "0.0.2"
  sha256 "0e2dd53c6109ce7df63e28c03be5a788fd12d85cc144d8a3c64b50b9b3443601"

  def install
    bin.install "chezmoi"
  end
end
