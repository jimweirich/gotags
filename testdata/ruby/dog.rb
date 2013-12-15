module Animal

  VERSION = "1.0"

  class Base
  end

  class Dog < Animal::Base
    attr_reader :tail, :head,
      :front_feet, :Back_Feet, , :Fur

    def speak(string)
    end
    alias :bark :speak

    def wag
    end
    alias_method :wiggle, :wag

    def EatFood
    end
    alias :Consume :Eat

    def _internal
    end
  end

  Mammal = Stuct.new(:head, :feet)
end
