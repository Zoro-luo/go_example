Git*： //https://www.liaoxuefeng.com/wiki/896043488029600/896067074338496
	(1) 在Linux上安装git*:
		sudo apt-get install git
	(2) git管理:
		a. 创建git版本库
			mkdir gitdir	cd gitdir
			git init	//创建版本库,git创建了一个master分支
		b. git常规指令用法：
			git status	
			git add . 	//文件修改添加到暂存区
			git commit -m "说明" 	//把暂存区的所有修改提交到当前分支
			git push origin master
		c. 版本回退：
			git log  或者  git log --pretty=oneline 	 //显示每一次commit的记录(含commitId)
			git log --pretty=oneline --abbrev-commit //短commitid 显示每次提交的记录
			git reset --hard HEAD^	//回退上一个版本
			git reset --hard 1094a  //回退到指定的commitId的版本
			git reflog	//显示 记录的每一次指令(含版本的commitID)
		d. 撤销修改：
			git checkout -- user.php //把user.php文件在工作区的修改全部撤销
		e. 删除文件:
			git rm text.txt 						
			git commit -m "remove test.txt"	//当git status 告知工作区中有文件被删除了 则git rm会删除版本中该文件
	(3) 远程仓库：
			git remote add origin git@github.com:michaelliao/learngit.git 	//本地仓库关联远程github仓库
			git push -u origin master	//本地的所有内容推送到远程origin仓库上
	(4) 从远程库克隆:
			git clone git@github.com:michaelliao/gitskills.git
	(5) 分支管理：
		a. 创建和合并分支：
			git checkout -b dev 	//创建了一个dev分支并切换到dev分支 -b参数表示创建并切换 相当于下面两条指令：git branch dev  git checkout dev
			git branch	//查看当前分支 命令会列出所有分支，当前分支前面会标一个*号。
			git add readme.txt  git commit -m "branch test"  //在dev分支上做修改并提交
			git checkout master		//切换回master分支(会发现在dev分支上修改的readme.txt文件没做修改)
			git merge dev	//把dev分支的工作成果合并到master分支上(会发现master分支和dev修改完全一样)
			git branch -d dev 	//完成合并后 就可以放心的删除dev分支
		b. Bug分支：
			git stash 	//把当前工作现场"储藏"起来，等以后恢复现场后继续工作
			git status	//被储藏后再用git status查看工作区，就是干净的.因此可以放心地创建分支来修复bug
			//首先确定要在哪个分支上修复bug，假定需要在master分支上修复，就从master创建临时分支
			git checkout master
			git checkout -b issue-101	//创建一个bug(issue-101)分支并切换
			git add readme.txt 	 git commit -m "fix bug 101" //修复bug的操作
			git checkout master	//修复完成后，切换到master分支
			git merge --no-ff -m "merged bug fix 101" issue-101	//完成合并
			git branch -d issue-101 	//完成合并后删除issue-101分支
			//太棒了，原计划两个小时的bug修复只花了5分钟！现在，是时候接着回到dev分支干活了！
			git checkout dev //回到dev分支继续干活
			git status 	//dev分区的工作区是干净的
			//工作区是干净的，刚才的工作现场存到哪去了？用git stash list命令看看：
			git stash list
			//工作现场还在，Git把stash内容存在某个地方了，但是需要恢复一下，有两个办法：
			//一是用git stash apply恢复，但是恢复后，stash内容并不删除，你需要用git stash drop来删除；
			//另一种方式是用git stash pop，恢复的同时把stash内容也删了
			git stash pop 
			git stash list 	//再用git stash list查看，就看不到任何stash内容了
		e. 多人协作：
			git remote -v  //查看远程库的详细信息
			git push origin dev //推送开发分支
		f. 标签管理：  (标签也是版本库的一个快照，标签总是和某个commit挂钩)
			git tag v1.0   //提交commit版本后 创建一个tag标签 (默认标签是打在最新提交的commit上的)
			git tag 	   //查看所有标签
			//有时候，如果忘了打标签，比如，现在已经是周五了，但应该在周一打的标签没有打，怎么办？
			//方法是找到历史提交的commit id，然后打上就可以了
			git log --pretty=oneline --abbrev-commit
			git tag v0.9 f52c633 //后面带上commit id 就是在此次commit提交后打上tag
			git show v0.9 		//查看标签信息
			git tag -d v0.1  	//删除tag
			//因为创建的标签都只存储在本地，不会自动推送到远程。所以，打错的标签可以在本地安全删除。
			//如果要推送某个标签到远程，使用命令git push origin <tagname>：
			git push origin v1.0
			git push origin --tags  	//一次性推送全部尚未推送到远程的本地标签
	(6) 自定义Git指令别名：	
			https://www.liaoxuefeng.com/wiki/896043488029600/900785521032192