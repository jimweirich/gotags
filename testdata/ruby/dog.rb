module Animal

  VERSION = "1.0"

  class Base
  end

  class Dog < Animal::Base
    attr_reader :tail, :head,
      :feet,                    # comment

    def speak(string)
    end
    alias :bark :speak

    def wag
    end
    alias_method :wiggle, :wag
  end
end
