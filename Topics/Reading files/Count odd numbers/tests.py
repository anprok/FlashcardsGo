from hstest import StageTest, CheckResult, WrongAnswer, TestCase

line_count = ["4999"]


class Test(StageTest):
    def generate(self):
        return [TestCase(stdin=[test], attach=test) for test in line_count]

    def check(self, reply, clue):
        if reply.strip() != clue.strip():
            return CheckResult.wrong("Wrong amount of odd numbers counted!\n"
                                     "Your program counted {} odd numbers, please check your code!".format(reply.strip()))

        return CheckResult.correct()


if __name__ == '__main__':
    Test('file.file').run_tests()