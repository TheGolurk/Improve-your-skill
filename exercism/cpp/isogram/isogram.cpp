#include "isogram.h"

#include <map>
#include <cctype>

namespace isogram {

    bool is_isogram(const std::string word) {
        std::map<char, int> mp;

        for (char character : word){
            if (!std::isalpha(character)) {
                continue;
            }

            character = std::toupper(character);
            mp[character] += 1;

            if (mp[character] > 1) {
                return false;
            }

        }

        return true;
    }

}  // namespace isogram
