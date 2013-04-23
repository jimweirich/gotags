module Animal

  VERSION = "1.0"

  class Base
  end

  class Dog < Animal::Base
    attr_reader :tail, :head

    def speak(string)
    end
    alias :bark :speak

    def wag
    end
  end
end
