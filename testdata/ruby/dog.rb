module Animal

  VERSION = "1.0"

  class Dog
    attr_reader :tail, :head

    def speak(string)
    end
    alias :bark :speak

    def wag
    end
  end
end
