class Kaalin < Formula
  desc "Qaraqalpaq tili ushın CLI vositası — latin↔kirill konvertatsiya, son→so'z"
  homepage "https://github.com/dontbeidle/kaalin"
  license "MIT"

  on_macos do
    on_arm do
      url "https://github.com/dontbeidle/kaalin/releases/download/v#{version}/kaalin_#{version}_darwin_arm64.tar.gz"
    end
    on_intel do
      url "https://github.com/dontbeidle/kaalin/releases/download/v#{version}/kaalin_#{version}_darwin_amd64.tar.gz"
    end
  end

  on_linux do
    on_arm do
      url "https://github.com/dontbeidle/kaalin/releases/download/v#{version}/kaalin_#{version}_linux_arm64.tar.gz"
    end
    on_intel do
      url "https://github.com/dontbeidle/kaalin/releases/download/v#{version}/kaalin_#{version}_linux_amd64.tar.gz"
    end
  end

  def install
    bin.install "kaalin"
    generate_completions_from_executable(bin/"kaalin", "completion")
  end

  test do
    system bin/"kaalin", "version"
  end
end
