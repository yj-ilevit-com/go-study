## mysql root 패스워드로 접근 안될 경우

1. mysql 서버 중지
   brew services stop mysql
2. 인증 우회
   /opt/homebrew/opt/mysql/bin/mysqld_safe --skip-grant-tables &
3. Root 접속
   mysql -u root

## source 경로설정도 pwd를 통해 확인
